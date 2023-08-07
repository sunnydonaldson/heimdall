use reqwest::Request;
use std::fmt::Debug;

pub trait LimitingRule: Debug {
    fn allows(&self, request: &Request) -> bool;
}
