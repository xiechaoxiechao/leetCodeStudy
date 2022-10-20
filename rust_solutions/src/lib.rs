
pub fn add(left: usize, right: usize) -> usize {
    left + right
}
pub mod rs1108;
mod rs1991;
mod solution;
mod  rs191;
mod  rs169;


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
        'a:loop{
            break 'a;
        }
        let a=[1;5];
    }
    
}
