package dto

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
	"go.arcalot.io/log"
)

type ACT struct {
	Revision         string `yaml:"revision"`
	Image_Name       string `default:"all"`
	Project_Filepath string
	Image_Tag        string `default:"latest"`
	Quay_Img_Exp     string `default:"never"`
	Build_Timeout    uint32 `default:"600"`
	Registries       []Registry
}

func Unmarshal(logger log.Logger) (ACT, error) {
	filteredRegistries, err := UnmarshalRegistries(logger)
	if err != nil {
		return ACT{}, err
	}
	conf := ACT{
		Revision:         viper.GetString("revision"),
		Image_Name:       viper.GetString("image_name"),
		Project_Filepath: viper.GetString("project_filepath"),
		Image_Tag:        viper.GetString("image_tag"),
		Quay_Img_Exp:     viper.GetString("quay_img_exp"),
		Build_Timeout:    viper.GetUint32("build_timeout"),
		Registries:       filteredRegistries}
	if err := defaults.Set(&conf); err != nil {
		return ACT{}, fmt.Errorf("error setting defaults (%w)", err)
	}
	return conf, nil
}
