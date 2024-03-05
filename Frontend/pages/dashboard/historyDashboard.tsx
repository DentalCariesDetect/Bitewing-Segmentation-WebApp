import { SideBar } from "@/components/SideBar";
import { MagnifyingGlassIcon } from "@heroicons/react/24/outline";
import { PencilIcon, UserPlusIcon } from "@heroicons/react/24/solid";
import {
    Card,
    CardHeader,
    Input,
    Typography,
    Button,
    CardBody,
    CardFooter,
    Avatar,
    IconButton,
    Tooltip,
} from "@material-tailwind/react";
import { motion } from 'framer-motion';
import { fadeIn } from '@/variants';
import Link from "next/link";


export default function historyDashboard() {



    const TABLE_HEAD = ["Patient ID", "Gender", "Age", "Phone", "Edit", ""];

    const TABLE_ROWS = [
        {
            img: "https://scontent.fbkk5-4.fna.fbcdn.net/v/t1.6435-9/86263526_2565482020399840_7973518111129206784_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=be3454&_nc_eui2=AeGteQNRBY3cll7biRo3CVIUwibETxdg7HHCJsRPF2DscSeSXHvDnbQihINj0OpCu4I4Ix4DnCLZhnIUMyGjVcJS&_nc_ohc=1GrZLN6aR3QAX_HHZcX&_nc_oc=AQn9YHotEmC3qt623G6lYmhXn1FJmttafoJ81Cs--XDRZqUEwMizZ2_XlbYBogbP_Ss&_nc_ht=scontent.fbkk5-4.fna&oh=00_AfCBy42us10tXn4pYhGN9f4xHKZB9EtwnFQNEbrQ3OTeMA&oe=65E9A183",
            name: "John Michael",
            email: "john@creative-tim.com",
            gender: "Male",
            org: "Organization",
            online: true,
            phonenumber: "0930000000",
        },
        {
            img: "https://scontent.fbkk5-4.fna.fbcdn.net/v/t1.6435-9/86263526_2565482020399840_7973518111129206784_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=be3454&_nc_eui2=AeGteQNRBY3cll7biRo3CVIUwibETxdg7HHCJsRPF2DscSeSXHvDnbQihINj0OpCu4I4Ix4DnCLZhnIUMyGjVcJS&_nc_ohc=1GrZLN6aR3QAX_HHZcX&_nc_oc=AQn9YHotEmC3qt623G6lYmhXn1FJmttafoJ81Cs--XDRZqUEwMizZ2_XlbYBogbP_Ss&_nc_ht=scontent.fbkk5-4.fna&oh=00_AfCBy42us10tXn4pYhGN9f4xHKZB9EtwnFQNEbrQ3OTeMA&oe=65E9A183",
            name: "Alexa Liras",
            email: "alexa@creative-tim.com",
            gender: "Female",
            org: "Developer",
            online: false,
            phonenumber: "0930000000",
        },
        {
            img: "https://scontent.fbkk5-4.fna.fbcdn.net/v/t1.6435-9/86263526_2565482020399840_7973518111129206784_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=be3454&_nc_eui2=AeGteQNRBY3cll7biRo3CVIUwibETxdg7HHCJsRPF2DscSeSXHvDnbQihINj0OpCu4I4Ix4DnCLZhnIUMyGjVcJS&_nc_ohc=1GrZLN6aR3QAX_HHZcX&_nc_oc=AQn9YHotEmC3qt623G6lYmhXn1FJmttafoJ81Cs--XDRZqUEwMizZ2_XlbYBogbP_Ss&_nc_ht=scontent.fbkk5-4.fna&oh=00_AfCBy42us10tXn4pYhGN9f4xHKZB9EtwnFQNEbrQ3OTeMA&oe=65E9A183",
            name: "Laurent Perrier",
            email: "laurent@creative-tim.com",
            gender: "Male",
            org: "Projects",
            online: false,
            phonenumber: "0930000000",
        },
        {
            img: "https://scontent.fbkk5-4.fna.fbcdn.net/v/t1.6435-9/86263526_2565482020399840_7973518111129206784_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=be3454&_nc_eui2=AeGteQNRBY3cll7biRo3CVIUwibETxdg7HHCJsRPF2DscSeSXHvDnbQihINj0OpCu4I4Ix4DnCLZhnIUMyGjVcJS&_nc_ohc=1GrZLN6aR3QAX_HHZcX&_nc_oc=AQn9YHotEmC3qt623G6lYmhXn1FJmttafoJ81Cs--XDRZqUEwMizZ2_XlbYBogbP_Ss&_nc_ht=scontent.fbkk5-4.fna&oh=00_AfCBy42us10tXn4pYhGN9f4xHKZB9EtwnFQNEbrQ3OTeMA&oe=65E9A183",
            name: "Michael Levi",
            email: "michael@creative-tim.com",
            gender: "Male",
            org: "Developer",
            online: true,
            phonenumber: "0930000000",
        },
        {
            img: "https://scontent.fbkk5-4.fna.fbcdn.net/v/t1.6435-9/86263526_2565482020399840_7973518111129206784_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=be3454&_nc_eui2=AeGteQNRBY3cll7biRo3CVIUwibETxdg7HHCJsRPF2DscSeSXHvDnbQihINj0OpCu4I4Ix4DnCLZhnIUMyGjVcJS&_nc_ohc=1GrZLN6aR3QAX_HHZcX&_nc_oc=AQn9YHotEmC3qt623G6lYmhXn1FJmttafoJ81Cs--XDRZqUEwMizZ2_XlbYBogbP_Ss&_nc_ht=scontent.fbkk5-4.fna&oh=00_AfCBy42us10tXn4pYhGN9f4xHKZB9EtwnFQNEbrQ3OTeMA&oe=65E9A183",
            name: "Richard Gran",
            email: "richard@creative-tim.com",
            gender: "Female",
            org: "Executive",
            online: false,
            phonenumber: "0930000000",
        },

    ];

    return (
        <div className="w-full h-screen">
            {/* <Transition /> */}

            <div
                className="bg-gradient-background h-full flex items-center justify-between ">

                <SideBar />

                <motion.div
                    variants={fadeIn('right', 0.05)}
                    initial="hidden"
                    animate="show"
                    exit="hidden"
                    className="h-full/8 w-full m-10 bg-blue-900 text-white p-3 rounded-xl">
                    <div className=" flex place-content-between">
                        <div className=" items-end">
                            <h1 className=" text-3xl mt-5 ml-5">Predict History</h1>
                        </div>
                        <div>
                            <Link href="../predictDentist">
                                <button className="bg-white text-black rounded-md p-2 mt-5 ml-5">Predict</button>
                            </Link>
                        </div>
                    </div>

                    <CardHeader floated={false} shadow={false} className="rounded-none">

                        <div className="flex flex-col items-center justify-between gap-4 md:flex-row bg-blue-900">

                        </div>
                    </CardHeader>
                    <CardBody className="overflow-scroll px-0">
                        <table className="mt-4 w-full min-w-max table-auto text-left">
                            <thead>
                                <tr>
                                    {TABLE_HEAD.map((head) => (
                                        <th
                                            key={head}
                                            className="border-y border-blue-gray-100 bg-blue-gray-50/50 p-4"
                                        >
                                            <Typography
                                                variant="small"
                                                color="blue-gray"
                                                className="font-normal leading-none opacity-70"
                                            >
                                                {head}
                                            </Typography>
                                        </th>
                                    ))}
                                </tr>
                            </thead>
                            <tbody>
                                {TABLE_ROWS.map(
                                    ({ img, name, email, gender, org, online, phonenumber }, index) => {
                                        const isLast = index === TABLE_ROWS.length - 1;
                                        const classes = isLast
                                            ? "p-4"
                                            : "p-4 border-b border-purple-200";

                                        return (
                                            <tr key={name}>
                                                <td className={classes}>
                                                    <div className="flex items-center gap-3">
                                                        <Avatar src={img} alt={name} size="sm" className="w-10 h-10"
                                                        />
                                                        <div className="flex flex-col">
                                                            <Typography
                                                                variant="small"
                                                                color="blue-gray"
                                                                className="font-normal"
                                                            >
                                                                {name}
                                                            </Typography>
                                                            <Typography
                                                                variant="small"
                                                                color="blue-gray"
                                                                className="font-normal opacity-70"
                                                            >
                                                                {email}
                                                            </Typography>
                                                        </div>
                                                    </div>
                                                </td>
                                                <td className={classes}>
                                                    <div className="flex flex-col">
                                                        <Typography
                                                            variant="small"
                                                            color="blue-gray"
                                                            className="font-normal"
                                                        >
                                                            {gender}
                                                        </Typography>
                                                        <Typography
                                                            variant="small"
                                                            color="blue-gray"
                                                            className="font-normal opacity-70"
                                                        >
                                                            {org}
                                                        </Typography>
                                                    </div>
                                                </td>
                                                <td className={classes}>
                                                    <div className="w-max">
                                                        31
                                                    </div>
                                                </td>
                                                <td className={classes}>
                                                    <Typography
                                                        variant="small"
                                                        color="blue-gray"
                                                        className="font-normal"
                                                    >
                                                        {phonenumber}
                                                    </Typography>
                                                </td>
                                                <td className={classes}>
                                                    <Tooltip content="Edit User">
                                                        <IconButton variant="text">
                                                            <PencilIcon className="h-4 w-4" />
                                                        </IconButton>
                                                    </Tooltip>
                                                </td>
                                            </tr>
                                        );
                                    },
                                )}
                            </tbody>
                        </table>
                    </CardBody>
                    <CardFooter className="flex items-center justify-between border-t border-blue-gray-50 p-4">
                        <Typography variant="small" color="blue-gray" className="font-normal">
                            Page 1 of 10
                        </Typography>
                        <div className="flex gap-2">
                            <Button variant="outlined" size="sm">
                                Previous
                            </Button>
                            <Button variant="outlined" size="sm">
                                Next
                            </Button>
                        </div>
                    </CardFooter>
                </motion.div>
            </div>
            <div></div>
        </div>
    );
}
