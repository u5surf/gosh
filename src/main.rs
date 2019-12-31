use std::io::{stdin,stdout,Write};
fn main() {
    loop {
        let mut command = String::new();
        print!("\x1b[0;32mgosh \x1b[0;34mλ \x1b[0m ");
        stdout().flush().unwrap();
        stdin().read_line(&mut command).unwrap();
    }
}
