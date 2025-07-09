import type { ChangeEvent } from 'react';

import { Button, InputText } from 'components';
import api from 'lib/api/api.ts';

import useData from './useData';

const SignUp = () => {
  const { errors, handleSubmit, navigate, onChange, values } = useData();

  const onSubmit = () =>
    handleSubmit(async (value) => {
      if (!value) return;
      const resp = await api.register(value);
      if (resp.status && resp.data?.access_token) {
        window.localStorage.setItem('access_token', resp.data.access_token);
        navigate('/', { replace: true });
      }
    });

  return (
    <div className="grid place-items-center gap-12 m-0 h-screen px-8 text-[#aaaaaa] bg-[#eff9ff] sm:p-0">
      <div className="fixed top-[-50vmin] left-[-50vmin] w-[100vmin] h-[100vmin] bg-[#65c8ff] rounded-[47%_53%_61%_39%/45%_51%_49%_55%] after:content-[''] after:fixed after:right-[-50vmin] after:bottom-[-55vmin] after:w-[100vmin] after:h-[100vmin] after:bg-[#143d81] after:rounded-[47%_53%_61%_39%/45%_51%_49%_55%]"></div>

      <div className="overflow-hidden relative z-[3] w-11/12 px-8 pt-44 pb-14 rounded-3xl bg-white text-center shadow-xl before:content-[''] before:absolute before:top-[-880px] before:left-1/2 before:-translate-x-1/2 before:w-[1000px] before:h-[1000px] before:rounded-[50%] before:bg-[#216ce7] before:z-[-1] sm:w-[500px]">
        <h2 className="text-white text-2xl font-inter-extra-bold mb-4 absolute top-8 left-1/2 -translate-x-1/2">
          Todo App
        </h2>
        <h2 className="text-2xl mb-10 text-black font-inter-bold">Sign Up</h2>
        <div className="grid gap-3 mb-11">
          <InputText
            type="text"
            placeholder="Enter your full name"
            name="name"
            id="name"
            label="Full Name"
            value={values.name}
            onChange={(e: ChangeEvent<HTMLInputElement>) => onChange({ name: e.target.value })}
            error={errors.name}
          />
          <InputText
            type="text"
            placeholder="Enter your email address"
            name="email"
            id="email"
            label="Email"
            value={values.email}
            onChange={(e: ChangeEvent<HTMLInputElement>) => onChange({ email: e.target.value })}
            error={errors.email}
          />
          <InputText
            type="password"
            placeholder="Enter your password"
            name="password"
            id="password"
            label="Password"
            value={values.password}
            onChange={(e: ChangeEvent<HTMLInputElement>) => onChange({ password: e.target.value })}
            error={errors.password}
          />
          <Button onClick={onSubmit}>
            <span className="text-[16px]">Create Account</span>
          </Button>
        </div>
        <footer>
          Already have an account? Sign in{' '}
          <Button variant="text" type="link" href="/sign-in" className="underline text-blue-700 text-[16px]">
            here
          </Button>
        </footer>
      </div>
    </div>
  );
};

export default SignUp;
