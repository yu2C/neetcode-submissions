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
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut map = HashMap::new();

        for (i, v) in inorder.iter().enumerate() {
            map.insert(*v, i);
        }

        fn helper(
            preorder: &Vec<i32>,
            map: &HashMap<i32, usize>,
            pre_l: usize,
            pre_r: usize,
            in_l: usize,
            in_r: usize
        ) -> Option<Rc<RefCell<TreeNode>>> {
            if pre_l > pre_r {
                return None;
            }

            let root_val = preorder[pre_l];
            let root_i = map[&root_val];
            let node = Rc::new(RefCell::new(TreeNode::new(root_val)));
            let l_size = root_i - in_l;

            node.borrow_mut().left = helper(
                preorder,
                map,
                pre_l + 1,
                pre_l + l_size,
                in_l,
                root_i - 1,
            );

            node.borrow_mut().right = helper(
                preorder,
                map,
                pre_l + l_size + 1,
                pre_r,
                root_i + 1,
                in_r
            );

            Some(node)
        }
        if preorder.is_empty() {
            return None;
        }
        helper(&preorder, &map, 0, preorder.len()-1, 0, inorder.len()-1)
    }
}
