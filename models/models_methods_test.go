package models

import (
	"strings"
	"testing"
	"time"

	"gopkg.in/yaml.v2"

	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
	"github.com/stretchr/testify/require"
)

// ----------------------------
// --- Validate

// Config
func TestValidateConfig(t *testing.T) {
	t.Log("Valid bitriseData ID")
	{
		bitriseData := BitriseDataModel{
			Workflows: map[string]WorkflowModel{
				"A-Za-z0-9-_.": WorkflowModel{},
			},
		}

		warnings, err := bitriseData.Validate()
		require.NoError(t, err)
		require.Equal(t, 0, len(warnings))
	}

	t.Log("Invalid bitriseData ID - empty")
	{
		bitriseData := BitriseDataModel{
			Workflows: map[string]WorkflowModel{
				"": WorkflowModel{},
			},
		}
		warnings, err := bitriseData.Validate()
		require.NoError(t, err)
		require.Equal(t, 1, len(warnings))
	}

	t.Log("Invalid bitriseData ID - contains: `/`")
	{
		bitriseData := BitriseDataModel{
			Workflows: map[string]WorkflowModel{
				"wf/id": WorkflowModel{},
			},
		}

		warnings, err := bitriseData.Validate()
		require.NoError(t, err)
		require.Equal(t, 1, len(warnings))
	}

	t.Log("Invalid bitriseData ID - contains: `:`")
	{
		bitriseData := BitriseDataModel{
			Workflows: map[string]WorkflowModel{
				"wf:id": WorkflowModel{},
			},
		}

		warnings, err := bitriseData.Validate()
		require.NoError(t, err)
		require.Equal(t, 1, len(warnings))
	}

	t.Log("Invalid bitriseData ID - contains: ` `")
	{
		bitriseData := BitriseDataModel{
			Workflows: map[string]WorkflowModel{
				"wf id": WorkflowModel{},
			},
		}

		warnings, err := bitriseData.Validate()
		require.NoError(t, err)
		require.Equal(t, 1, len(warnings))
	}
}

// Workflow
func TestValidateWorkflow(t *testing.T) {
	t.Log("before-afetr test")
	{
		workflow := WorkflowModel{
			BeforeRun: []string{"befor1", "befor2", "befor3"},
			AfterRun:  []string{"after1", "after2", "after3"},
		}

		warnings, err := workflow.Validate()
		require.NoError(t, err)
		require.Equal(t, 0, len(warnings))
	}

	t.Log("invalid workflow - Invalid env: more than 2 fields")
	{
		configStr := `
format_version: 1.0.0
default_step_lib_source: "https://github.com/bitrise-io/bitrise-steplib.git"

workflows:
  target:
    envs:
    - ENV_KEY: env_value
      opts:
        title: test_env
    title: Output Test
    steps:
    - script:
        title: Should fail
        inputs:
        - content: echo "Hello"
          BAD_KEY: value
`

		config := BitriseDataModel{}
		require.NoError(t, yaml.Unmarshal([]byte(configStr), &config))
		require.NoError(t, config.Normalize())

		warnings, err := config.Validate()
		require.Error(t, err)
		require.Equal(t, true, strings.Contains(err.Error(), "Invalid env: more than 2 fields:"))
		require.Equal(t, 0, len(warnings))
	}

	t.Log("vali workflow - Warning: duplicated inputs")
	{
		configStr := `
format_version: 1.0.0
default_step_lib_source: "https://github.com/bitrise-io/bitrise-steplib.git"

workflows:
  target:
    steps:
    - script:
        title: Should fail
        inputs:
        - content: echo "Hello"
        - content: echo "Hello"
`

		config := BitriseDataModel{}
		require.NoError(t, yaml.Unmarshal([]byte(configStr), &config))
		require.NoError(t, config.Normalize())

		warnings, err := config.Validate()
		require.NoError(t, err)
		require.Equal(t, 1, len(warnings))
	}
}

// ----------------------------
// --- Merge

func TestMergeEnvironmentWith(t *testing.T) {
	diffEnv := envmanModels.EnvironmentItemModel{
		"test_key": "test_value",
		envmanModels.OptionsKey: envmanModels.EnvironmentItemOptionsModel{
			Title:             pointers.NewStringPtr("test_title"),
			Description:       pointers.NewStringPtr("test_description"),
			Summary:           pointers.NewStringPtr("test_summary"),
			ValueOptions:      []string{"test_valu_options1", "test_valu_options2"},
			IsRequired:        pointers.NewBoolPtr(true),
			IsExpand:          pointers.NewBoolPtr(false),
			IsDontChangeValue: pointers.NewBoolPtr(true),
			IsTemplate:        pointers.NewBoolPtr(true),
		},
	}

	t.Log("Different keys")
	{
		env := envmanModels.EnvironmentItemModel{
			"test_key1": "test_value",
		}
		require.Error(t, MergeEnvironmentWith(&env, diffEnv))
	}

	t.Log("Normal merge")
	{
		env := envmanModels.EnvironmentItemModel{
			"test_key":              "test_value",
			envmanModels.OptionsKey: envmanModels.EnvironmentItemOptionsModel{},
		}
		require.NoError(t, MergeEnvironmentWith(&env, diffEnv))

		options, err := env.GetOptions()
		require.NoError(t, err)

		diffOptions, err := diffEnv.GetOptions()
		require.NoError(t, err)

		require.Equal(t, *diffOptions.Title, *options.Title)
		require.Equal(t, *diffOptions.Description, *options.Description)
		require.Equal(t, *diffOptions.Summary, *options.Summary)
		require.Equal(t, len(diffOptions.ValueOptions), len(options.ValueOptions))
		require.Equal(t, *diffOptions.IsRequired, *options.IsRequired)
		require.Equal(t, *diffOptions.IsExpand, *options.IsExpand)
		require.Equal(t, *diffOptions.IsDontChangeValue, *options.IsDontChangeValue)
		require.Equal(t, *diffOptions.IsTemplate, *options.IsTemplate)
	}
}

func TestMergeStepWith(t *testing.T) {
	desc := "desc 1"
	summ := "sum 1"
	website := "web/1"
	fork := "fork/1"
	published := time.Date(2012, time.January, 1, 0, 0, 0, 0, time.UTC)

	stepData := stepmanModels.StepModel{
		Description:         pointers.NewStringPtr(desc),
		Summary:             pointers.NewStringPtr(summ),
		Website:             pointers.NewStringPtr(website),
		SourceCodeURL:       pointers.NewStringPtr(fork),
		PublishedAt:         pointers.NewTimePtr(published),
		HostOsTags:          []string{"osx"},
		ProjectTypeTags:     []string{"ios"},
		TypeTags:            []string{"test"},
		IsRequiresAdminUser: pointers.NewBoolPtr(true),
		Inputs: []envmanModels.EnvironmentItemModel{
			envmanModels.EnvironmentItemModel{
				"KEY_1": "Value 1",
			},
			envmanModels.EnvironmentItemModel{
				"KEY_2": "Value 2",
			},
		},
		Outputs: []envmanModels.EnvironmentItemModel{},
	}

	diffTitle := "name 2"
	newSuppURL := "supp"
	runIfStr := ""
	stepDiffToMerge := stepmanModels.StepModel{
		Title:      pointers.NewStringPtr(diffTitle),
		HostOsTags: []string{"linux"},
		Source: stepmanModels.StepSourceModel{
			Git: "https://git.url",
		},
		Dependencies: []stepmanModels.DependencyModel{
			stepmanModels.DependencyModel{
				Manager: "brew",
				Name:    "test",
			},
		},
		SupportURL: pointers.NewStringPtr(newSuppURL),
		RunIf:      pointers.NewStringPtr(runIfStr),
		Inputs: []envmanModels.EnvironmentItemModel{
			envmanModels.EnvironmentItemModel{
				"KEY_2": "Value 2 CHANGED",
			},
		},
	}

	mergedStepData, err := MergeStepWith(stepData, stepDiffToMerge)
	require.NoError(t, err)

	require.Equal(t, "name 2", *mergedStepData.Title)
	require.Equal(t, "desc 1", *mergedStepData.Description)
	require.Equal(t, "sum 1", *mergedStepData.Summary)
	require.Equal(t, "web/1", *mergedStepData.Website)
	require.Equal(t, "fork/1", *mergedStepData.SourceCodeURL)
	require.Equal(t, true, (*mergedStepData.PublishedAt).Equal(time.Date(2012, time.January, 1, 0, 0, 0, 0, time.UTC)))
	require.Equal(t, "linux", mergedStepData.HostOsTags[0])
	require.Equal(t, "", *mergedStepData.RunIf)
	require.Equal(t, 1, len(mergedStepData.Dependencies))

	dep := mergedStepData.Dependencies[0]
	require.Equal(t, "brew", dep.Manager)
	require.Equal(t, "test", dep.Name)

	// inputs
	input0 := mergedStepData.Inputs[0]
	key0, value0, err := input0.GetKeyValuePair()

	require.NoError(t, err)
	require.Equal(t, "KEY_1", key0)
	require.Equal(t, "Value 1", value0)

	input1 := mergedStepData.Inputs[1]
	key1, value1, err := input1.GetKeyValuePair()

	require.NoError(t, err)
	require.Equal(t, "KEY_2", key1)
	require.Equal(t, "Value 2 CHANGED", value1)
}

func TestGetInputByKey(t *testing.T) {
	stepData := stepmanModels.StepModel{
		Inputs: []envmanModels.EnvironmentItemModel{
			envmanModels.EnvironmentItemModel{
				"KEY_1": "Value 1",
			},
			envmanModels.EnvironmentItemModel{
				"KEY_2": "Value 2",
			},
		},
	}

	_, found := getInputByKey(stepData, "KEY_1")
	require.Equal(t, true, found)

	_, found = getInputByKey(stepData, "KEY_3")
	require.Equal(t, false, found)
}

// ----------------------------
// --- StepIDData

func TestGetStepIDStepDataPair(t *testing.T) {
	stepData := stepmanModels.StepModel{}

	t.Log("valid steplist item")
	{
		stepListItem := StepListItemModel{
			"step1": stepData,
		}

		id, _, err := GetStepIDStepDataPair(stepListItem)
		require.NoError(t, err)
		require.Equal(t, "step1", id)
	}

	t.Log("invalid steplist item - more than 1 step")
	{
		stepListItem := StepListItemModel{
			"step1": stepData,
			"step2": stepData,
		}

		id, _, err := GetStepIDStepDataPair(stepListItem)
		require.Error(t, err)
		require.Equal(t, "", id)
	}
}

func TestCreateStepIDDataFromString(t *testing.T) {
	t.Log("default / long / verbose ID mode")
	{
		stepCompositeIDString := "steplib-src::step-id@0.0.1"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.NoError(t, err)
		require.Equal(t, "steplib-src", stepIDData.SteplibSource)
		require.Equal(t, "step-id", stepIDData.IDorURI)
		require.Equal(t, "0.0.1", stepIDData.Version)
	}

	t.Log("no steplib-source")
	{
		stepCompositeIDString := "step-id@0.0.1"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "default-steplib-src")

		require.NoError(t, err)
		require.Equal(t, "default-steplib-src", stepIDData.SteplibSource)
		require.Equal(t, "step-id", stepIDData.IDorURI)
		require.Equal(t, "0.0.1", stepIDData.Version)
	}

	t.Log("invalid/empty step lib source, but default provided")
	{
		stepCompositeIDString := "::step-id@0.0.1"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "default-steplib-src")

		require.NoError(t, err)
		require.Equal(t, "default-steplib-src", stepIDData.SteplibSource)
		require.Equal(t, "step-id", stepIDData.IDorURI)
		require.Equal(t, "0.0.1", stepIDData.Version)
	}

	t.Log("invalid/empty step lib source + no default")
	{
		stepCompositeIDString := "::step-id@0.0.1"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.Error(t, err)
		require.Equal(t, "", stepIDData.SteplibSource)
		require.Equal(t, "", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	t.Log("no steplib-source & no default -> fail")
	{
		stepCompositeIDString := "step-id@0.0.1"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.Error(t, err)
		require.Equal(t, "", stepIDData.SteplibSource)
		require.Equal(t, "", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	t.Log("no steplib & no version, only step-id")
	{
		stepCompositeIDString := "step-id"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "def-lib-src")

		require.NoError(t, err)
		require.Equal(t, "def-lib-src", stepIDData.SteplibSource)
		require.Equal(t, "step-id", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	t.Log("empty test")
	{
		stepCompositeIDString := ""
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "def-step-src")

		require.Error(t, err)
		require.Equal(t, "", stepIDData.SteplibSource)
		require.Equal(t, "", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	t.Log("special empty test")
	{
		stepCompositeIDString := "@1.0.0"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "def-step-src")

		require.Error(t, err)
		require.Equal(t, "", stepIDData.SteplibSource)
		require.Equal(t, "", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	//
	// ----- Local Path
	t.Log("local Path")
	{
		stepCompositeIDString := "path::/some/path"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.NoError(t, err)
		require.Equal(t, "path", stepIDData.SteplibSource)
		require.Equal(t, "/some/path", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	t.Log("local Path")
	{
		stepCompositeIDString := "path::~/some/path/in/home"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.NoError(t, err)
		require.Equal(t, "path", stepIDData.SteplibSource)
		require.Equal(t, "~/some/path/in/home", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	t.Log("local Path")
	{
		stepCompositeIDString := "path::$HOME/some/path/in/home"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.NoError(t, err)
		require.Equal(t, "path", stepIDData.SteplibSource)
		require.Equal(t, "$HOME/some/path/in/home", stepIDData.IDorURI)
		require.Equal(t, "", stepIDData.Version)
	}

	//
	// ----- Direct git uri
	t.Log("direct git uri")
	{
		stepCompositeIDString := "git::https://github.com/bitrise-io/steps-timestamp.git@develop"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "some-def-coll")

		require.NoError(t, err)
		require.Equal(t, "git", stepIDData.SteplibSource)
		require.Equal(t, "https://github.com/bitrise-io/steps-timestamp.git", stepIDData.IDorURI)
		require.Equal(t, "develop", stepIDData.Version)
	}

	t.Log("direct git uri")
	{
		stepCompositeIDString := "git::git@github.com:bitrise-io/steps-timestamp.git@develop"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.NoError(t, err)
		require.Equal(t, "git", stepIDData.SteplibSource)
		require.Equal(t, "git@github.com:bitrise-io/steps-timestamp.git", stepIDData.IDorURI)
		require.Equal(t, "develop", stepIDData.Version)
	}

	//
	// ----- Old step
	t.Log("old step")
	{
		stepCompositeIDString := "_::https://github.com/bitrise-io/steps-timestamp.git@1.0.0"
		stepIDData, err := CreateStepIDDataFromString(stepCompositeIDString, "")

		require.NoError(t, err)
		require.Equal(t, "_", stepIDData.SteplibSource)
		require.Equal(t, "https://github.com/bitrise-io/steps-timestamp.git", stepIDData.IDorURI)
		require.Equal(t, "1.0.0", stepIDData.Version)
	}
}

// ----------------------------
// --- RemoveRedundantFields

func TestRemoveEnvironmentRedundantFields(t *testing.T) {
	t.Log("Trivial remove - all fields should be default value")
	{
		env := envmanModels.EnvironmentItemModel{
			"TEST_KEY": "test_value",
			envmanModels.OptionsKey: envmanModels.EnvironmentItemOptionsModel{
				Title:             pointers.NewStringPtr(""),
				Description:       pointers.NewStringPtr(""),
				Summary:           pointers.NewStringPtr(""),
				ValueOptions:      []string{},
				IsRequired:        pointers.NewBoolPtr(envmanModels.DefaultIsRequired),
				IsExpand:          pointers.NewBoolPtr(envmanModels.DefaultIsExpand),
				IsDontChangeValue: pointers.NewBoolPtr(envmanModels.DefaultIsDontChangeValue),
				IsTemplate:        pointers.NewBoolPtr(envmanModels.DefaultIsTemplate),
			},
		}
		require.NoError(t, removeEnvironmentRedundantFields(&env))

		options, err := env.GetOptions()
		require.NoError(t, err)

		require.Equal(t, (*string)(nil), options.Title)
		require.Equal(t, (*string)(nil), options.Description)
		require.Equal(t, (*string)(nil), options.Summary)
		require.Equal(t, 0, len(options.ValueOptions))
		require.Equal(t, (*bool)(nil), options.IsRequired)
		require.Equal(t, (*bool)(nil), options.IsExpand)
		require.Equal(t, (*bool)(nil), options.IsDontChangeValue)
		require.Equal(t, (*bool)(nil), options.IsTemplate)
	}

	t.Log("Trivial don't remove - no fields should be default value")
	{
		env := envmanModels.EnvironmentItemModel{
			"TEST_KEY": "test_value",
			envmanModels.OptionsKey: envmanModels.EnvironmentItemOptionsModel{
				Title:             pointers.NewStringPtr("t"),
				Description:       pointers.NewStringPtr("d"),
				Summary:           pointers.NewStringPtr("s"),
				ValueOptions:      []string{"i"},
				IsRequired:        pointers.NewBoolPtr(true),
				IsExpand:          pointers.NewBoolPtr(false),
				IsDontChangeValue: pointers.NewBoolPtr(true),
				IsTemplate:        pointers.NewBoolPtr(true),
			},
		}
		require.NoError(t, removeEnvironmentRedundantFields(&env))

		options, err := env.GetOptions()
		require.NoError(t, err)

		require.Equal(t, "t", *options.Title)
		require.Equal(t, "d", *options.Description)
		require.Equal(t, "s", *options.Summary)
		require.Equal(t, "i", options.ValueOptions[0])
		require.Equal(t, true, *options.IsRequired)
		require.Equal(t, false, *options.IsExpand)
		require.Equal(t, true, *options.IsDontChangeValue)
		require.Equal(t, true, *options.IsTemplate)
	}

	t.Log("No options - opts field shouldn't exist")
	{
		env := envmanModels.EnvironmentItemModel{
			"TEST_KEY": "test_value",
		}
		require.NoError(t, removeEnvironmentRedundantFields(&env))

		_, ok := env[envmanModels.OptionsKey]
		require.Equal(t, false, ok)
	}
}

func configModelFromYAMLBytes(configBytes []byte) (bitriseData BitriseDataModel, err error) {
	if err = yaml.Unmarshal(configBytes, &bitriseData); err != nil {
		return
	}
	return
}

func TestRemoveWorkflowRedundantFields(t *testing.T) {
	configStr := `
format_version: 1.0.0
default_step_lib_source: "https://github.com/bitrise-io/bitrise-steplib.git"

app:
  summary: "sum"
  envs:
  - ENV_KEY: env_value
    opts:
      is_required: true

workflows:
  target:
    envs:
    - ENV_KEY: env_value
      opts:
        title: test_env
    title: Output Test
    steps:
    - script:
        description: test
`

	config, err := configModelFromYAMLBytes([]byte(configStr))
	require.NoError(t, err)

	err = config.RemoveRedundantFields()
	require.NoError(t, err)

	require.Equal(t, "", config.App.Title)
	require.Equal(t, "", config.App.Description)
	require.Equal(t, "sum", config.App.Summary)

	for _, env := range config.App.Environments {
		options, err := env.GetOptions()
		require.NoError(t, err)

		require.Nil(t, options.Title)
		require.Nil(t, options.Description)
		require.Nil(t, options.Summary)
		require.Equal(t, 0, len(options.ValueOptions))
		require.Equal(t, true, *options.IsRequired)
		require.Nil(t, options.IsExpand)
		require.Nil(t, options.IsDontChangeValue)
	}

	for _, workflow := range config.Workflows {
		require.Equal(t, "Output Test", workflow.Title)
		require.Equal(t, "", workflow.Description)
		require.Equal(t, "", workflow.Summary)

		for _, env := range workflow.Environments {
			options, err := env.GetOptions()
			require.NoError(t, err)

			require.Equal(t, "test_env", *options.Title)
			require.Nil(t, options.Description)
			require.Nil(t, options.Summary)
			require.Equal(t, 0, len(options.ValueOptions))
			require.Nil(t, options.IsRequired)
			require.Nil(t, options.IsExpand)
			require.Nil(t, options.IsDontChangeValue)
		}

		for _, stepListItem := range workflow.Steps {
			_, step, err := GetStepIDStepDataPair(stepListItem)
			require.NoError(t, err)

			require.Nil(t, step.Title)
			require.Equal(t, "test", *step.Description)
			require.Nil(t, step.Summary)
			require.Nil(t, step.Website)
			require.Nil(t, step.SourceCodeURL)
			require.Nil(t, step.SupportURL)
			require.Nil(t, step.PublishedAt)
			require.Equal(t, "", step.Source.Git)
			require.Equal(t, "", step.Source.Commit)
			require.Equal(t, 0, len(step.HostOsTags))
			require.Equal(t, 0, len(step.ProjectTypeTags))
			require.Equal(t, 0, len(step.TypeTags))
			require.Nil(t, step.IsRequiresAdminUser)
			require.Nil(t, step.IsAlwaysRun)
			require.Nil(t, step.IsSkippable)
			require.Nil(t, step.RunIf)
			require.Equal(t, 0, len(step.Inputs))
			require.Equal(t, 0, len(step.Outputs))
		}
	}
}
