use super::config::Config;
use reqwest::Request;

#[derive(Debug)]
pub struct Heimdall {
    config: Config,
}

impl Heimdall {
    pub fn new(config: Config) -> Self {
        Self { config }
    }

    pub fn process(&self, request: &Request) {
        let actions = if self
            .config
            .limiting_rules
            .iter()
            .all(|rule| rule.allows(request))
        {
            &self.config.failure_actions
        } else {
            &self.config.failure_actions
        };
        actions.iter().for_each(|action| action.execute(request));
    }
}
