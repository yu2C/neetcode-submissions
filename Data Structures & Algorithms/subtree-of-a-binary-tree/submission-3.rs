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
    pub fn is_subtree(root: Option<Rc<RefCell<TreeNode>>>, sub_root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        fn same_tree(a: Option<Rc<RefCell<TreeNode>>>, b: Option<Rc<RefCell<TreeNode>>>) -> bool {
            match (a, b) {
                (None, None) => true,
                (None, Some(_)) | (Some(_), None) => false,
                (Some(a), Some(b)) => {
                    let a = a.borrow();
                    let b = b.borrow();

                    a.val == b.val
                    && same_tree(a.right.clone(), b.right.clone()) 
                    && same_tree(a.left.clone(), b.left.clone())
                }
            }
        }
        match (root, sub_root) {
            (_, None) => true,
            (None, Some(_)) => false,
            (Some(root), Some(sub_root)) => {
                same_tree(Some(root.clone()), Some(sub_root.clone())) 
                || Self::is_subtree(root.borrow().left.clone(), Some(sub_root.clone()))
                || Self::is_subtree(root.borrow().right.clone(), Some(sub_root))
            }
        }
    }
}
