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

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn is_balanced(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        fn dfs(node: Option<Rc<RefCell<TreeNode>>>) -> (bool, i32) {
            match node {
                None => (true, 0),

                Some(node) => {
                    let node = node.borrow();
                    let (l_b, l_h) = dfs(node.left.clone());
                    let (r_b, r_h) = dfs(node.right.clone());

                    let b = l_b && r_b && (l_h - r_h).abs() <= 1;
                    let h = l_h.max(r_h)+1;
                    (b, h)
                }
            }
        }
        let (b, _) = dfs(root);
        return b
    }
}
