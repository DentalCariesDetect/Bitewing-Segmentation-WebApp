import React from "react";
import {
    Card,
    Typography,
    List,
    ListItem,
    ListItemPrefix,
    Alert,
} from "@material-tailwind/react";
import {
    UserCircleIcon,
    Cog6ToothIcon,
    PowerIcon,
    HomeIcon,
} from "@heroicons/react/24/solid";
import {
    CubeTransparentIcon,
} from "@heroicons/react/24/outline";
import Link from "next/link";
import Modal from "./Modal";

export function SideBar() {
    const [open, setOpen] = React.useState(0);
    const [openAlert, setOpenAlert] = React.useState(true);
    // logout popup state
    const [openLogout, setOpenLogout] = React.useState(false);

    const handleOpen = (value: any) => {
        setOpen(open === value ? 0 : value);
    };

    return (
        <Card className="h-full w-full bg-violet-800 max-w-[20rem] p-4 shadow-xl shadow-blue-gray-900/5 rounded-l-none border-white">
            <div className="mb-2 p-4">
                <Typography variant="h5" color="blue-gray" className=" text-gray-200 ">
                    Dashboard
                </Typography>
            </div>
            <List className=" space-y-3">
                {/* <Accordion
                    open={open === 1}
                    icon={
                        <ChevronDownIcon
                            strokeWidth={2.5}
                            className={`mx-auto h-4 w-4 transition-transform ${open === 1 ? "rotate-180" : ""}`}
                        />
                    }
                >
                    <ListItem className="p-0" selected={open === 1}>
                        <AccordionHeader onClick={() => handleOpen(1)} className="border-b-0 p-3">
                            <ListItemPrefix>
                                <PresentationChartBarIcon className="h-5 w-5" />
                            </ListItemPrefix>
                            <Typography color="blue-gray" className="mr-auto font-normal">
                                Dashboard
                            </Typography>
                        </AccordionHeader>
                    </ListItem>
                    <AccordionBody className="py-1">
                        <List className="p-0">
                            <ListItem>
                                <ListItemPrefix>
                                    <ChevronRightIcon strokeWidth={3} className="h-3 w-5" />
                                </ListItemPrefix>
                                Analytics
                            </ListItem>
                            <ListItem>
                                <ListItemPrefix>
                                    <ChevronRightIcon strokeWidth={3} className="h-3 w-5" />
                                </ListItemPrefix>
                                Reporting
                            </ListItem>
                            <ListItem>
                                <ListItemPrefix>
                                    <ChevronRightIcon strokeWidth={3} className="h-3 w-5" />
                                </ListItemPrefix>
                                Projects
                            </ListItem>
                        </List>
                    </AccordionBody>
                </Accordion> */}
                {/* <Accordion
                    open={open === 2}
                    icon={
                        <ChevronDownIcon
                            strokeWidth={2.5}
                            className={`mx-auto h-4 w-4 transition-transform ${open === 2 ? "rotate-180" : ""}`}
                        />
                    }
                >
                    <ListItem className="p-0" selected={open === 2}>
                        <AccordionHeader onClick={() => handleOpen(2)} className="border-b-0 p-3">
                            <ListItemPrefix>
                                <ShoppingBagIcon className="h-5 w-5" />
                            </ListItemPrefix>
                            <Typography color="blue-gray" className="mr-auto font-normal">
                                E-Commerce
                            </Typography>
                        </AccordionHeader>
                    </ListItem>
                    <AccordionBody className="py-1">
                        <List className="p-0">
                            <ListItem>
                                <ListItemPrefix>
                                    <ChevronRightIcon strokeWidth={3} className="h-3 w-5" />
                                </ListItemPrefix>
                                Orders
                            </ListItem>
                            <ListItem>
                                <ListItemPrefix>
                                    <ChevronRightIcon strokeWidth={3} className="h-3 w-5" />
                                </ListItemPrefix>
                                Products
                            </ListItem>
                        </List>
                    </AccordionBody>
                </Accordion> */}
                {/* <hr className="my-2 border-blue-gray-50" /> */}
                {/* Home */}
                <Link href="../dashboard">
                    <ListItem className="text-gray-100 bg-blue-900">
                        <ListItemPrefix className=" mr-2">
                            <HomeIcon className="h-5 w-5" />
                        </ListItemPrefix>
                        Home
                    </ListItem>
                </Link>
                {/* Profile Setting */}
                <Link href="../dashboard/profileSetting">
                    <ListItem className="text-gray-100 bg-blue-900 ">
                        <ListItemPrefix className=" mr-2">
                            <UserCircleIcon className="h-5 w-5" />
                        </ListItemPrefix>
                        Profile Setting
                    </ListItem>
                </Link>

                {/* Pateint History */}
                <Link href="../dashboard/historyDashboard">
                    <ListItem className="text-gray-100 bg-blue-900 ">
                        <ListItemPrefix className=" mr-2">
                            <Cog6ToothIcon className="h-5 w-5" />
                        </ListItemPrefix>
                        Pateint History
                    </ListItem>
                </Link>

                <button onClick={() => setOpenLogout(true)}>
                    <ListItem className="text-gray-100 bg-blue-900 ">
                        <ListItemPrefix className=" mr-2">
                            <PowerIcon className="h-5 w-5" />
                        </ListItemPrefix>
                        Log Out
                    </ListItem>
                </button>

                {/* if openLogout is true show modal component */}
                <Modal
                    isOpen={openLogout}
                    setIsOpen={setOpenLogout}
                    title="Log Out"
                    message="Are you sure you want to log out?"
                />

            </List>
            {/* <Alert open={openAlert} className="mt-auto bg-white text-black" onClose={() => setOpenAlert(false)}>
                <CubeTransparentIcon className="mb-4 h-12 w-12" />
                <Typography variant="h6" className="mb-1">
                    Upgrade to PRO
                </Typography>
                <Typography variant="small" className="font-normal opacity-80">
                    Upgrade to Material Tailwind PRO and get even more components, plugins, advanced features
                    and premium.
                </Typography>
                <div className="mt-4 flex gap-3">
                    <Typography
                        as="a"
                        href="#"
                        variant="small"
                        className="font-medium opacity-80"
                        onClick={() => setOpenAlert(false)}
                    >
                        Dismiss
                    </Typography>
                    <Typography as="a" href="#" variant="small" className="font-medium">
                        Upgrade Now
                    </Typography>
                </div>
            </Alert> */}
        </Card>
    );
}
