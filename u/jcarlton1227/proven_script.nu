use std assert

# Nushell
# A new type of shell
def main [
    no_default: string,
    name = "Nicolas Bourbaki",
    age: int = 42,
    date_of_birth?: datetime,
    obj: record = {"records": "included"},
    l: list<string> = ["or", "lists!"],
    tables?: table,
    enable_kill_mode?: bool = true,
] {
    # Test
    # https://www.nushell.sh/book/testing.html
		assert ($age == 42)

    print $"Hello World and a warm welcome especially to ($name)"
    print "and its acolytes.." $age $obj $l
    print $tables

    let secret = try { 
      get_variable f/examples/secret
    } catch { 
      'No secret yet at f/examples/secret !' 
    };

    print $"The variable at `f/examples/secret`: ($secret)"
    # fetch context variables
    let user = $env.WM_USERNAME

    # Nu pipelines
    ls | where size > 1kb | sort-by modified | print "ls:" $in

    # Nu works with existing data
    # Nu speaks JSON, YAML, SQLite, Excel, and more out of the box. 
    # It's easy to bring data into a Nu pipeline whether it's in a file, a database, or a web API:
    let nu_license = http get https://api.github.com/repos/nushell/nushell | get license

    return { splitted: ($name | split words), user: $user, nu_license: $nu_license}
    # Interested in learning more?
    # https://www.nushell.sh/book/getting_started.html
}
