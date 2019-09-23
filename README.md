# Dad Joke Terraform Provider

That's right: a Terraform provider that let's you declare dad jokes as resources.

![Feels Good](https://i.kym-cdn.com/entries/icons/original/000/006/077/so_good.png)

I feel the same way.

## Why The Weird Name?

This provider interfaces with the [icanhazdadjoke](https://icanhazdadjoke.com/) API, hence the "ichdj" at the end of the provider name ("**I** **C**an **H**az **D**ad **J**oke").

## Usage

**NOTE**: Compiled only for Linux thus far.

1. Obtain the `terraform-provider-ichdj` binary from the [releases page](https://github.com/aidanSoles/terraform-provider-ichdj/releases). You can also unzip the `terraform-provider-ichdj-<TARGET_OS>.zip` file to get the provider binary.
2. Place the provider binary in the same directory as your dad joke resource declarations (`terraform` will look in your current directory for missing providers). **NOTE**: If you don't want to have a binary sitting next to your `*.tf` files, a more elegant solution can be found [here](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).
3. Declare your desired dad joke resources (basic example shown below, and in [main.tf](https://github.com/aidanSoles/terraform-provider-ichdj/blob/master/main.tf)), then `terraform init`, `terraform plan`, and `terraform apply`!

```
resource "ichdj_random_joke" "first" {
  name = "first"
}

output "joke" {
  value = "${ichdj_random_joke.first.joke}"
}
```

## Compilation

To compile the provider yourself, [set up Go on your local system](https://golang.org/doc/install), clone this repo, and run:

```
go build -o terraform-provider-ichdj
```

in the project's root directory.
