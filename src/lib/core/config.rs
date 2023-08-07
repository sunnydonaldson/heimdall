use std::fmt::Debug;

use super::action::Action;
use super::limiting_rule::LimitingRule;

#[derive(Debug)]
pub struct Config {
    pub limiting_rules: Vec<Box<dyn LimitingRule>>,
    pub success_actions: Vec<Box<dyn Action>>,
    pub failure_actions: Vec<Box<dyn Action>>,
}

impl Config {
    pub fn new(
        limiting_rules: Vec<Box<dyn LimitingRule>>,
        success_actions: Vec<Box<dyn Action>>,
        failure_actions: Vec<Box<dyn Action>>,
    ) -> Self {
        Config {
            limiting_rules,
            success_actions,
            failure_actions,
        }
    }
}
