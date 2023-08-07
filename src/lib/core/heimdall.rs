use super::config::Config;

#[derive(Debug)]
pub struct Heimdall {
    config: Config,
}

impl Heimdall {
    pub fn new(config: Config) -> Self {
        Self { config }
    }
}
