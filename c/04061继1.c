struct TreeNode* inorderSuccessor(struct TreeNode* root, struct TreeNode* p) {
    struct TreeNode *successor = NULL;
    if (p->right != NULL) {
        successor = p->right;
        while (successor->left != NULL) {
            successor = successor->left;
        }
        return successor;
    }
    struct TreeNode *node = root;
    while (node != NULL) {
        if (node->val > p->val) {
            successor = node;
            node = node->left;
        } else {
            node = node->right;
        }
    }
    return successor;
}


