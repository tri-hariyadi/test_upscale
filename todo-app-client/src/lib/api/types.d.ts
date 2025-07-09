export interface IApiRequestParams {
  url: string;
  payload?: ILooseObject;
  isMultipart?: boolean;
  isErrorless?: boolean;
}

export interface IApi<T = Any> {
  Response: {
    status: boolean;
    message: string;
    data?: T;
  };

  AuthResponse: {
    access_token: string;
  };

  LoginRequest: {
    email: string;
    password: string;
  };

  RegisterRequest: {
    name: string;
    email: string;
    password: string;
  };

  UserProfile: {
    id: string;
    name: string;
    email: string;
  };

  TodosRequest: {
    title: string;
    description: string;
    tag: string;
    due_date: string;
  };

  Todos: {
    id: number;
    title: string;
    description: string;
    tag: string;
    due_date: string;
    created_at: string;
    updated_at: string;
  };

  Quote: {
    quote: string;
    author: string;
    category: string;
  };
}
