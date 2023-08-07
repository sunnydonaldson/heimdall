mod core;
use self::core::config::Config;
use self::core::heimdall::Heimdall;

pub fn start_heimdall() {
    let heimdall = Heimdall::new(Config::new(vec![], vec![]));
    println!("heimdall constructed");
    println!("{:?}", heimdall);
}
