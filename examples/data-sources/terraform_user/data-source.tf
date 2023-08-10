data "terraform_user" "my_user" {
    id = 123
            password = "k!5As3HquUrQ"
            root_role = {
        integer = 1
    }
            send_email = false
        }