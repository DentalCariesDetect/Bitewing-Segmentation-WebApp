declare module "next-connect" {
  interface NextConnectOptions {
    // Define any options you need, for instance:
    onError?: (err: Error, req: any, res: any, next: any) => void;
    onNoMatch?: (req: any, res: any) => void;
    // Add more options as per your requirement
  }

  function NextConnect(options?: NextConnectOptions): any; // Replace 'any' with a more specific type if needed
  export = NextConnect;
}
