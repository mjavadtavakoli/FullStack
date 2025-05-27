import { useState } from 'react';
import { register } from '../services/auth';

export default function Register() {
  const [form, setForm] = useState({ name: '', email: '', password: '' });
  const [msg, setMsg] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const result = await register(form);
    setMsg(result.message || result.error);
  };

  return (
    <div>
      <h2>ثبت‌نام</h2>
      <form onSubmit={handleSubmit}>
        <input placeholder="نام" onChange={e => setForm({ ...form, name: e.target.value })} />
        <input placeholder="ایمیل" onChange={e => setForm({ ...form, email: e.target.value })} />
        <input placeholder="رمز عبور" type="password" onChange={e => setForm({ ...form, password: e.target.value })} />
        <button type="submit">ثبت‌نام</button>
      </form>
      <p>{msg}</p>
    </div>
  );
}
