import { z } from "zod";
import { PostSchema } from "./post.validation";

export type BasicPostProps = z.infer<typeof PostSchema> & { slug: string };
