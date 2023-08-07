mod core;
use self::core::config::Config;
use self::core::heimdall::Heimdall;
use reqwest::{Method, Request, Url};

pub fn start_heimdall() {
    let heimdall = Heimdall::new(Config::new(vec![], vec![], vec![]));

    println!("heimdall constructed");
    println!("{:?}", heimdall);
    let request = Request::new(Method::GET, Url::parse("https://google.com").unwrap());
    heimdall.process(&request);
}
