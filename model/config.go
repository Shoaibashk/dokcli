package model

type Config struct {
	name    string                 `yaml: name`
	port    string                 `yaml: port`
	specUrl map[string]interface{} `yaml: spec-url`
}

// var Defaults = map[string]interface{}{
// 	"headertitle": "Dokcli",
// 	"tabtitle":    "dokcli",
// 	"specurl":     "https://petstore.swagger.io/v2/swagger.json",
// 	"name":        "Paper",
// 	"license":     "MIT",
// 	"licenselink": "https://github.com/nanxiaobei/hugo-paper/blob/main/LICENSE",
// 	"description": "A simple, clean, flexible Hugo theme",
// 	"homepage":    "https://github.com/nanxiaobei/hugo-paper/",
// 	"demosite":    "https://hugo-paper.vercel.app",
// 	"tags":        []string{"Responsive", "Simple", "Clean", "Light", "White", "Blog"},
// 	"features":    []string{"Responsive", "One Column", "Blog"},
// 	"min_version": "0.57.1",
// }

var Defaults = map[string]interface{}{
	"name":     "Docuctl",
	"port":     "1221",
	"spec-url": map[string]interface{}{"petstore": "https://petstore.swagger.io/v2/swagger.json"},
}
