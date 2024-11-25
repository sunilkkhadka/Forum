import { Link } from "react-router-dom";

const Post = () => {
  return (
    <section>
      <h1>Posts:</h1>
      <Link to="/post/create">Create Post</Link>
    </section>
  );
};

export default Post;
