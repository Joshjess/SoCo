import { Button, Form, Input } from 'antd';
import axios from 'axios';
import { redirect } from "react-router-dom";
import { useNavigate } from "react-router-dom";

  


const headers = {
  'Content-Type': 'text/plain'
};

// const goToLogin = () => {
  //   return redirect("/login");
  // };
  
  const SignUp = () => {
    
  const navigate = useNavigate();
  
  const onFinish = (values) => {
    console.log('Success:', values);
    // console.log('Matching', values.password === values.rePassword)

    axios.post('http://localhost:8080/v1/users/signup', {
      email: values.username,
      password: values.password,
    },
    {headers}
    )
    .then(function (response) {
      console.log(response);
      navigate("/login");
    })
    .catch(function (error) {
      console.log(error);
    });
  };
  
  const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
  };
  
  return (
    <Form
      name="basic"
      labelCol={{
        span: 8,
      }}
      wrapperCol={{
        span: 16,
      }}
      initialValues={{
        remember: false,
      }}
      // validateTrigger
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[
          {
            required: true,
            message: 'Please input your username!',
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[
          {
            required: true,
            message: 'Please input your password!',
          },
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item
        label="Retype Password"
        name="rePassword"
        rules={[
          {
            required: true,
            message: 'Please repeat your password!',
          },
          ({ getFieldValue }) => ({
            validator(_, value) {
              if (!value || getFieldValue('password') === value) {
                return Promise.resolve();
              }
              return Promise.reject(new Error('The two passwords that you entered do not match!'));
            },
          }),
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item
        wrapperCol={{
          offset: 8,
          span: 16,
        }}
      >
        <Button type="primary" htmlType="submit">
          SignUp
        </Button>
      </Form.Item>
    </Form>
  );
};

export default SignUp;
