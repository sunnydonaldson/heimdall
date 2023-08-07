use reqwest::Request;
use std::fmt::Debug;

pub trait Action: Debug {
    fn execute(&self, request: &Request);
}
