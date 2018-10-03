package main

import (
  "encoding/json"
  "fmt"
  "flag"
  "os"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/secretsmanager"
)

func getSecretManager() *secretsmanager.SecretsManager {
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))
  return secretsmanager.New(sess)
}

func getSecretInput(key string) *secretsmanager.GetSecretValueInput {
  input := &secretsmanager.GetSecretValueInput{
      SecretId: aws.String(key),
  }
  return input
}

func handleError(err error) {
  if aerr, ok := err.(awserr.Error); ok {
      switch aerr.Code() {
      case secretsmanager.ErrCodeResourceNotFoundException:
          fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
      case secretsmanager.ErrCodeInvalidParameterException:
          fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
      case secretsmanager.ErrCodeInvalidRequestException:
          fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
      case secretsmanager.ErrCodeDecryptionFailure:
          fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
      case secretsmanager.ErrCodeInternalServiceError:
          fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
      default:
          fmt.Println(aerr.Error())
      }
  } else {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      fmt.Println(err.Error())
  }
  return
}

func main() {
  secret_key := flag.String("secret", "", "Secret Key in Secret Manager")
  is_json := flag.Bool("json", false, "Parse as JSON and assume {\"ENVIRONMENT_VARIABLE\": \"Value\"}")
  flag.Parse()
  svc := getSecretManager()
  input := getSecretInput(*secret_key)
  result, err := svc.GetSecretValue(input)
  if err != nil {
    handleError(err)
    os.Exit(-1)
  }
  if(*is_json) {
    var json_data map[string]interface{}
    err := json.Unmarshal([]byte(*result.SecretString), &json_data)
    if(err == nil) {
      for k, v := range json_data {
        fmt.Printf("export %s=%s\n",k,v)
      }
    }
  } else {
    fmt.Println(*result.SecretString)
  }
}
