import { SubmitHandler, useForm } from "react-hook-form";

import { BasicPostProps } from "../post.type";
import { client } from "../../../shared/api/http.api";

const PostForm = () => {
  const { register, handleSubmit } = useForm<BasicPostProps>();

  const onCreatePost: SubmitHandler<BasicPostProps> = (data) => {
    console.log(data);
    const fetchData = async () => {
      const response = await client.get("/post/all");
      console.log(response);
    };
    fetchData();
  };

  return (
    <section>
      <h1>Post Form</h1>
      <form onSubmit={handleSubmit(onCreatePost)}>
        <div>
          <input
            {...register("title")}
            type="text"
            name="title"
            placeholder="title"
          />
        </div>
        <div>
          <input
            {...register("description")}
            type="text"
            name="description"
            placeholder="description"
          />
        </div>
        <div>
          <textarea
            {...register("content")}
            placeholder="content"
            name="content"
          ></textarea>
        </div>
        <button type="submit">Create Post</button>
      </form>
    </section>
  );
};

export default PostForm;
