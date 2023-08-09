resource "unleash_user" "my_user" {
    email = "user@example.com"
            name = "User"
            password = "k!5As3HquUrQ"
            root_role = {
        integer = 1
    }
            send_email = false
            username = "hunter"
        }