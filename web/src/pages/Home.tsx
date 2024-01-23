import { useSessionProvider } from "@/service/SessionProvider";

export const HomePage = () => {
  const { email, role } = useSessionProvider();

  return (
    <div>
      <div>Logged in as: {email}</div>
      <div>Role: {role}</div>
    </div>
  );
};
