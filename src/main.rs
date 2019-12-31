use std::io::{stdin, stdout, Write};
use subprocess::{Exec,Redirection};
fn main() {
    loop {
        let mut command = String::new();
        print!("\x1b[0;32mgosh \x1b[0;34mλ \x1b[0m ");
        stdout().flush().unwrap();
        stdin().read_line(&mut command).unwrap();
        let command = Exec::cmd(command.trim())
            .stdout(Redirection::Pipe)
            .capture();
        let out = match command {
            Ok(n) => n.stdout_str(),
            Err(_) => {
                println!("Error");
                continue;
            },
        };
        println!("{}", out);
    }
}
