use std::fmt::Debug;

use super::action::Action;

#[derive(Debug)]
pub struct Config {
    success_actions: Vec<Box<dyn Action>>,
    failure_actions: Vec<Box<dyn Action>>,
}

impl Config {
    pub fn new(
        success_actions: Vec<Box<dyn Action>>,
        failure_actions: Vec<Box<dyn Action>>,
    ) -> Self {
        Config {
            success_actions,
            failure_actions,
        }
    }
}
