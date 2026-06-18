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
    pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
		fn dfs(node: Option<Rc<RefCell<TreeNode>>>) {
			if let Some(node) = node {
				let left;
				let right;

				{
					let mut nb = node.borrow_mut();

					left = nb.left.take();
					right = nb.right.take();

					nb.left = right.clone();
					nb.right = left.clone();
				}

				dfs(left);
				dfs(right);
			}
		}

		dfs(root.clone());
		return root
		
    }
}
