// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//     pub val: i32,
//     pub left: Option<Rc<RefCell<TreeNode>>>,
//     pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//     #[inline]
//     pub fn new(val: i32) -> Self {
//         TreeNode {
//             val,
//             left: None,
//             right: None,
//         }
//     }
// }

use std::cell::RefCell;
use std::rc::Rc;

struct Codec {}

impl Codec {
    fn new() -> Self {
        Codec{}
    }

    fn serialize(&self, root: Option<Rc<RefCell<TreeNode>>>) -> String {
        fn dfs(node: Option<Rc<RefCell<TreeNode>>>, out: &mut Vec<String>) {
            if node.is_none() {
                out.push("null".to_string());
                return
            }

            let n = node.unwrap();
            let nb = n.borrow();
            out.push(nb.val.to_string());
            dfs(nb.left.clone(), out);
            dfs(nb.right.clone(), out);
        }

        let mut out = Vec::new();
        dfs(root, &mut out);
        out.join(",")
    }

    fn deserialize(&self, data: String) -> Option<Rc<RefCell<TreeNode>>> {
        let tokens: Vec<&str> = data.split(',').collect();
        let mut i = 0usize;
        fn build(tokens: &Vec<&str>, i: &mut usize) -> Option<Rc<RefCell<TreeNode>>> {
            if tokens[*i] == "null" {
                *i += 1;
                return None;
            }
            let val = tokens[*i].parse::<i32>().unwrap();
            *i += 1;

            let node = Rc::new(RefCell::new(TreeNode::new(val)));
            let left = build(tokens, i);
            let right = build(tokens, i);

            node.borrow_mut().left = left;
            node.borrow_mut().right = right;
            Some(node)
        }
        build(&tokens, &mut i)
    }
}
