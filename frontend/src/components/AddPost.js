import { Button, Form, Input } from 'antd';
import axios from 'axios';
import { useCookies } from 'react-cookie';

const { TextArea } = Input;


function postPost(title, text, cookie) {
  console.log(title, text, cookie.token)

  let data = {
    Title: title,
    Text: text,
  };

  let headers = {
    'Content-Type': 'text/plain',
    'Authorization': 'Bearer ' + cookie.token,
  };

  axios.post('http://localhost:8080/v1/posts/create', data, {headers: headers})
  .catch(function (error) {
    console.log(error);
  })
  .then(function (response) {
    console.log(response);
  })
}


// function sendPost(title, text) {
//   axios.post('http://localhost:8080/v1/posts/create', 
//     {
//       Title: title, 
//       Text: text,
//     }, 
//     {headers}
//   )
//   .catch(function (error) {
//     console.log(error);
//   })
//   .then(function (response) {
//     console.log(response);
//   })
// }

const onFinish = (cookies) => (values) => {
    console.log('Success:', values);
    console.log(cookies)
    postPost(values.title, values.text, cookies)
}

function submitPost(title, text) {
  console.log('Submit:', title)
  console.log('Submit:', text);
}

function AddPost() {
  const [cookies, setCookie] = useCookies(['token'])



  return (
    <>
      <Form
          labelCol={{
            span: 4,
          }}
          wrapperCol={{
            span: 14,
          }}
          layout="horizontal"
          style={{
            maxWidth: 600,
          }}
          onFinish={onFinish(cookies)}
          // onFinishFailed={onFinishFailed}
      >
        <Form.Item
            name="title"
            rules={[
              {
                required: true,
                message: 'Please input your title!',
              },
            ]}
        >
          <Input placeholder="Title" showCount maxLength={20}/>
        </Form.Item>

        <Form.Item
            name="text"
            rules={[
              {
                required: true,
                message: 'Please input your content!',
              },
            ]}
        >
          <TextArea placeholder="Content" showCount maxLength={100}/>
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">Post</Button>
        </Form.Item>
      </Form>
    </>
)};

export default AddPost;
