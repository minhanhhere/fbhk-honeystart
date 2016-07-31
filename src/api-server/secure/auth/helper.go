package auth

func containsInSlice(list []interface{}, el interface{}) bool {
    for _, it := range list {
        if it == el {
            return true
        }
    }
    return false
}
