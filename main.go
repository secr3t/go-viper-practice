package main

import (
    "fmt"
    "github.com/spf13/viper"
    "log"
)


func main() {

    viper.SetEnvPrefix("TEST")
    viper.BindEnv("PROFILE")
    profile := viper.Get("PROFILE").(string)
    viper.SetConfigFile(fmt.Sprintf("configs/%s.yml", profile))

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(viper.Get("profile"))
    log.Println(viper.Get("logging.level"))
    log.Printf("%v", viper.AllSettings())

}