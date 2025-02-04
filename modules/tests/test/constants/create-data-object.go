package constants

const (
	ModifyDataObjectTaskName                     = "modify-data-object"
	ModifyDataObjectServiceAccountName           = "modify-data-object-task"
	ModifyDataObjectServiceAccountNameNamespaced = "modify-data-object-test"
	ModifyDataObjectTaskrunName                  = "taskrun-modify-data-object"

	UnusualRestartCountThreshold = 3
	ReasonError                  = "Error"
)

type modifyDataObjectParams struct {
	Manifest            string
	WaitForSuccess      string
	AllowReplace        string
	DeleteObject        string
	DeleteObjectName    string
	DeleteObjectKind    string
	DataObjectNamespace string
}

var ModifyDataObjectParams = modifyDataObjectParams{
	Manifest:            "manifest",
	WaitForSuccess:      "waitForSuccess",
	AllowReplace:        "allowReplace",
	DeleteObject:        "deleteObject",
	DeleteObjectName:    "deleteObjectName",
	DeleteObjectKind:    "deleteObjectKind",
	DataObjectNamespace: "namespace",
}

type modifyDataObjectResults struct {
	Name      string
	Namespace string
}

var ModifyDataObjectResults = modifyDataObjectResults{
	Name:      "name",
	Namespace: "namespace",
}
