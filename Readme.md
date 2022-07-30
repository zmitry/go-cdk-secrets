## how to start
go run .

### golang cdk example with google secret manager

According to clean architecture you might have several layers of your app , We consider secrets and configuration to be the outer layer of your app (framework and devices)

Outer layer of our app is main function, that's why we put secrets parsing there. 

![](https://miro.medium.com/max/1400/1*0u-ekVHFu7Om7Z-VTwFHvg.png)

secrets definitions come from Args struct that serves several purposes
1. it defines struct that we use during app setup
2. it generates help for our cli
3. it defines mapping for env variables to struct keys
4. it defines mapping for secret storage to our struct keys
5. we use args struct in test in order to pass parameters for local testing.
```
type Args struct {
	// key where secret is stored
	PostgresURLKey string `env:"POSTGRES_URL_KEY" long:"postgres-url-key" default:"postgres-url"`
	// actual url for current service
    // we can keep it empty and it will be inferred from 
    // postgres-url secret manager
    // or you can provide it from env or arguments and pass secret value directly.
	PostgresURL string `env:"POSTGRES_URL" long:"postgres-url" secret-key:"PostgresURLKey" optional:"true"`
}
```

```
`SECRET_BASE_URL` - helps to simplify definition of url keys, instead of --postgres-url-key=`gcpsecretmanager://projects/es-scalability-test/secrets/my-key` 
you can specify just --postgres-url-key=mykey but it's totally optional and you can pass full url of any of they keys.


here some example, you might even pass keys from different storages

--postgres-url-key=gcpsecretmanager://projects/es-scalability-test/secrets/my-key
--dynamo-db-url-key=awssecretsmanager://secret-variable-name?region=us-east-2&decoder=string

--my-meta-file=blob://myvar.txt?decoder=string, along with BLOBVAR_BUCKET_URL it will allow you to read data into argument from any bucket

this way you can pass local parameter 
--my-arg-key=constant://?val=hello+world&decoder=string
or this way which also will work.
--my-arg=hello-wrold

```