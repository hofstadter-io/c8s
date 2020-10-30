package schema

type mapping struct {
	// val for printing
	Val string

	// key for import
	Key string

	// directory path containing definition
	Dir string

	// Cue path to definition in file
	Path string
}

// maps import strings to import key
var imports map[string]string = map[string]string {
	"k8s.io/api/apps/v1": "appsV1",
}

var mappings map[string]mapping = map[string]mapping {

	"admissionreview": mapping{
		"admissionreview", "admissionV1", "k8s.io/api/admission/v1", "#AdmissionReview",
	},
	"admissionrequest": mapping{
		"admissionrequest", "admissionV1", "k8s.io/api/admission/v1", "#AdmissionRequest",
	},
	"admissionresponse": mapping{
		"admissionresponse", "admissionV1", "k8s.io/api/admission/v1", "#AdmissionResponse",
	},

	"deploy": mapping{
		"deployment", "appsV1", "k8s.io/api/apps/v1", "#DeploymentSpec",
	},
	"deployment": mapping{
		"deployment", "appsV1", "k8s.io/api/apps/v1", "#DeploymentSpec",
	},



	"#TypeMeta": mapping{
		"#TypeMeta", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#TypeMeta",
	},
	"#ListMeta": mapping{
		"#ListMeta", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#ListMeta",
	},
	"#ObjectMeta": mapping{
		"#ObjectMeta", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#ObjectMeta",
	},
	"#OwnerReference": mapping{
		"#OwnerReference", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#OwnerReference",
	},

	"#GetOptions": mapping{
		"#GetOptions", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#GetOptions",
	},
	"#ListOptions": mapping{
		"#ListOptions", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#ListOptions",
	},
	"#DeleteOptions": mapping{
		"#DeleteOptions", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#DeleteOptions",
	},
	"#CreateOptions": mapping{
		"#CreateOptions", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#CreateOptions",
	},
	"#PatchOptions": mapping{
		"#PatchOptions", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#PatchOptions",
	},

	"#Preconditions": mapping{
		"#Preconditions", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#Preconditions",
	},
	"#Status": mapping{
		"#Status", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#Status",
	},
	"#StatusDetails": mapping{
		"#StatusDetails", "metaV1", "k8s.io/apimachinery/pkg/apis/meta/v1", "#StatusDetails",
	},

}
