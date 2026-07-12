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
    pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut ans = i32::MIN;
        fn dfs(node: Option<Rc<RefCell<TreeNode>>>, ans: &mut i32) -> i32 {
            if node.is_none() {
                return 0;
            }

            let node_ref = node.unwrap();
            let node_borrow = node_ref.borrow();

            let left = dfs(node_borrow.left.clone(), ans);
            let right = dfs(node_borrow.right.clone(), ans);

            let left = left.max(0);
            let right = right.max(0);

            *ans = (*ans).max(node_borrow.val + left + right);
            node_borrow.val + left.max(right)

        }
        dfs(root, &mut ans);
        ans
    }
}
