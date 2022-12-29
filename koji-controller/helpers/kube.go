package helpers

import (
    "fmt"
    "k8s.io/apimachinery/pkg/types"
)

// returns the string representation of a types.NamespacedName
// $NAMESPACE/$NAME or just $NAME in case $NAMESPACE is empty
func KeyToString(key types.NamespacedName) string {
    name := key.Name
    namespace := key.Namespace

    if namespace == "" {
        return name
    }

    return fmt.Sprintf("%v/%v", namespace, name)
}
