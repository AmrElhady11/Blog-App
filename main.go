package main

import "blog/cmd"

func main() {
	cmd.Execute()
}

//version 1 we made before using cobra
//
//
//
//
//func serve() {
//	// installation gin -> go get -u github.com/gin-gonic/gin
//	// installation viber -> go get github.com/spf13/viper
//	//installation cobra -> go get -u github.com/spf13/cobra@latest
//
//	configs := configSet() // the func which we have implemented under
//
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message":  "pong",
//			"app name": viper.Get("app.name"),
//		})
//	})
//	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
//}
//
///*
//* we made this func after installation viber to make it easy
//* for app to config app name and server port and easy to edit
//* in one place
// */
//func configSet() config.Config {
//	viper.SetConfigName("config") // name of config file (without extension) and it is in config package
//	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
//	viper.AddConfigPath("config") // path to look for the config file in
//	if err := viper.ReadInConfig(); err != nil {
//		log.Fatal("Error reading config file")
//	}
//	var configs config.Config //  config is class in config package and Config is the struct
//	err := viper.Unmarshal(&configs)
//	if err != nil {
//		log.Fatal("unable to decode into struct")
//	}
//	return configs
//}
