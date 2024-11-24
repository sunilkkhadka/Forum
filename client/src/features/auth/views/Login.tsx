import { useDispatch, useSelector } from "react-redux";
import { zodResolver } from "@hookform/resolvers/zod";
import { SubmitHandler, useForm } from "react-hook-form";

import { loginUser } from "../auth.slice";
import { AppDispatch, RootState } from "../../../store/store";
import { IBasicFormFieldProps } from "../auth.type";
import { BasicFormSchema } from "../auth.validation";
import { toast } from "react-toastify";
import { useEffect } from "react";

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<IBasicFormFieldProps>({
    resolver: zodResolver(BasicFormSchema),
  });

  const dispatch = useDispatch<AppDispatch>();

  const { error, status } = useSelector((state: RootState) => state.auth);

  useEffect(() => {
    if (status == "succeeded") {
      toast.success("Logged In Successfully");
    } else if (status == "failed" && error) {
      toast.error(error.message);
    }
  }, [status, error]);

  const onLogin: SubmitHandler<IBasicFormFieldProps> = (data) => {
    dispatch(loginUser(data));
  };

  return (
    <section>
      <form onSubmit={handleSubmit(onLogin)}>
        <div>
          <input {...register("email")} type="text" placeholder="email" />
          {errors.email && <div>{errors.email.message}</div>}
        </div>
        <div>
          <input
            {...register("password")}
            type="password"
            placeholder="password"
          />
          {errors.password && <div>{errors.password.message}</div>}
        </div>
        <button type="submit">Log in</button>
      </form>
    </section>
  );
};

export default Login;
