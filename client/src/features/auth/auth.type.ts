import { z } from "zod";

import { BasicFormSchema, RegisterFormSchema } from "./auth.validation";

export type IBasicFormFieldProps = z.infer<typeof BasicFormSchema>;
export type IRegisterFormFieldProps = z.infer<typeof RegisterFormSchema>;
