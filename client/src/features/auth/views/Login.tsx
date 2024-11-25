import { useEffect } from "react";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom";
import { zodResolver } from "@hookform/resolvers/zod";
import { useDispatch, useSelector } from "react-redux";
import { SubmitHandler, useForm } from "react-hook-form";

import { loginUser } from "../auth.slice";
import { IBasicFormFieldProps } from "../auth.type";
import { BasicFormSchema } from "../auth.validation";
import { AppDispatch, RootState } from "../../../store/store";

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<IBasicFormFieldProps>({
    resolver: zodResolver(BasicFormSchema),
  });

  const navigate = useNavigate();

  const dispatch = useDispatch<AppDispatch>();

  const { error, status } = useSelector((state: RootState) => state.auth);

  useEffect(() => {
    if (status == "succeeded") {
      toast.success("Logged In Successfully");
      navigate("/post");
    } else if (status == "failed" && error) {
      toast.error(error.message);
    }
  }, [status, error, navigate]);

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
