import Link from "next/link";
import { headers, cookies } from "next/headers";
import { redirect } from "next/navigation";

export default function Login() {

    return (
        <div className="w-full h-screen">

            <div className="bg-gradient-background h-full flex items-center justify-center">
                <div className=" w-[500px] h-[600px] rounded-lg ">
                    <div className="flex justify-center items-center h-[100px]">
                        <h1 className="text-4xl font-bold">Login</h1>
                    </div>
                    <form
                        className="animate-in flex-1 flex flex-col w-full justify-center gap-2 text-foreground"

                    >
                        <label className="text-md" htmlFor="email">
                            Email
                        </label>
                        <input
                            className="rounded-md px-4 py-2 bg-inherit border mb-6"
                            name="email"
                            placeholder="Your email"
                            required
                        />
                        <label className="text-md" htmlFor="password">
                            Password
                        </label>
                        <input
                            className="rounded-md px-4 py-2 bg-inherit border mb-6"
                            type="password"
                            name="password"
                            placeholder="Your password"
                            required
                        />
                        <button className="bg-white text-black rounded-md px-4 py-2 text-foreground mb-2">
                            Login
                        </button>
                        <button
                            className="border border-foreground/20 rounded-md px-4 py-2 text-foreground mb-2"
                        >
                            Register
                        </button>

                    </form>

                </div>
            </div>
        </div>

    );
}
