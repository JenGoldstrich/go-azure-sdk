
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupprotecteditemscrr` Documentation

The `backupprotecteditemscrr` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-01-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupprotecteditemscrr"
```


### Client Initialization

```go
client := backupprotecteditemscrr.NewBackupProtectedItemsCrrClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupProtectedItemsCrrClient.BackupProtectedItemsCrrList`

```go
ctx := context.TODO()
id := backupprotecteditemscrr.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.BackupProtectedItemsCrrList(ctx, id, backupprotecteditemscrr.DefaultBackupProtectedItemsCrrListOperationOptions())` can be used to do batched pagination
items, err := client.BackupProtectedItemsCrrListComplete(ctx, id, backupprotecteditemscrr.DefaultBackupProtectedItemsCrrListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
