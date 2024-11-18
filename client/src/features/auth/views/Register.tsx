import { zodResolver } from "@hookform/resolvers/zod";
import { SubmitHandler, useForm } from "react-hook-form";

import { useRegisterUser } from "../hooks/useAuth";
import { IRegisterFormFieldProps } from "../auth.type";
import { RegisterFormSchema } from "../auth.validation";

const Register = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<IRegisterFormFieldProps>({
    resolver: zodResolver(RegisterFormSchema),
  });

  const registerMutation = useRegisterUser();

  const onSubmit: SubmitHandler<IRegisterFormFieldProps> = (data) => {
    console.log(data);
    registerMutation.mutate({
      email: data.email,
      password: data.password,
    });
  };

  return (
    <section>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <input {...register("email")} type="email" placeholder="email" />
          {errors.email && <p>{errors.email.message}</p>}
        </div>
        <div>
          <input
            {...register("password")}
            type="password"
            placeholder="password"
          />
          {errors.password && <p>{errors.password.message}</p>}
        </div>
        <div>
          <input
            {...register("confirmPassword")}
            type="password"
            placeholder="confirm password"
          />
          {errors.confirmPassword && <p>{errors.confirmPassword.message}</p>}
        </div>
        <button type="submit">Submit</button>
      </form>
    </section>
  );
};

export default Register;
