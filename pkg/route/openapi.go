package route

import (
	// restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "XGoPkg",
			Description: "Resource for managing package",
			Contact: &spec.ContactInfo{
				Name:  "xgopkg",
				Email: "dev_support@xgopkg.com",
				URL:   "http://xgopkg.com",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{{
		TagProps: spec.TagProps{
			Name:        "users",
			Description: "Managing users",
		},
	},
		{
			TagProps: spec.TagProps{
				Name:        "packages",
				Description: "Managing packages",
			},
		},
	}
}
