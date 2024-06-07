import json
import requests

URL = "http://localhost:9354/query"


def post_check(post_id):
    q = """
    query Post {{
        post(id: \"{}\") {{
            id
            title
            content
            author
            commentsAllowed
        }}
    }}""".format(post_id)
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        post = json.loads(r.text)["data"]["post"]
        if post["id"] == post_id:
            print("post check: passed")
            return
    
    print("post check: FAILED", r.text)


def posts_check(post_ids):
    q = """
    query Posts {
        posts {
            id
            title
            content
            author
            commentsAllowed
        }
    }
    """
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        posts = json.loads(r.text)["data"]["posts"]
        if len(posts) == len(post_ids):
            if post_ids == [i["id"] for i in posts]:
                print("posts check: passed")
                return
                
    print("posts check: FAILED")


def create_post_with_comments():
    q = """
    mutation CreatePost {
        createPost(
            title: "post_1"
            content: "post_1_content"
            author: "post_1_author"
            commentsAllowed: true
        ) {
        id
        title
        content
        author
        commentsAllowed
        }
    }
    """
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        print("create_post_with_comments: passed")
        
    return json.loads(r.text)["data"]["createPost"]["id"]


def create_post_without_comments():
    q = """
    mutation CreatePost {
        createPost(
            title: "post_1"
            content: "post_1_content"
            author: "post_1_author"
            commentsAllowed: false
        ) {
        id
        title
        content
        author
        commentsAllowed
        }
    }
    """
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        print("create_post_without_comments: passed")
        
    return json.loads(r.text)["data"]["createPost"]["id"]


def create_comment_root(post_id):
    q = """
    mutation CreateComment {{
        createComment(postId: "{}", author: "{}", content: "{}") {{
            id
            postId
            parentId
            author
            content
            createdAt
        }}
    }}
    """.format(
        post_id,
        "aleksei",
        "regular_comment"
    )
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        print("create_comment_root: passed")
        
    return json.loads(r.text)["data"]["createComment"]["id"]


def create_comment_child(post_id, root_comment_id):
    q = """
    mutation CreateComment {{
        createComment(
            postId:   "{}"
            author:   "{}"
            content:  "{}"
            parentId: "{}"
        ) {{
            id
            postId
            parentId
            author
            content
            createdAt
        }}
    }}
    """.format(
        post_id,
        "max",
        "big",
        root_comment_id,
    )
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        print("create_comment_child: passed")
        
    return json.loads(r.text)["data"]["createComment"]["id"]


def comments_check(comment_ids, post_id):
    q = """
    query Comments {{
        comments(postId: "{}", limit: {}, offset: {}) {{
            id
            postId
            parentId
            author
            content
            createdAt
        }}
    }}
    """.format(
        post_id,
        3,
        0
    )
    
    r = requests.post(
        URL,
        json={"query": q},
    )
    
    if r.status_code == 200:
        comments = json.loads(r.text)["data"]["comments"]
        if [i["id"] for i in comments] == comment_ids:
            print("comments check: passed")
            return
    
    print("comments check: FAILED")


def main():
    post1_id = create_post_with_comments()
    post2_id = create_post_without_comments()
    
    post_check(post1_id)
    posts_check([post1_id, post2_id])
    
    comment1_id = create_comment_root(post1_id)
    comment2_id = create_comment_child(post1_id, comment1_id)
    
    comments_check([comment1_id, comment2_id], post1_id)


if __name__ == "__main__":
    main()
