import Navbar from "@/components/Navbar";
import SupportedConfig from "@/components/SupportedConfig";
import Image from "next/image";
import Link from "next/link";
import { DiTerminal } from "react-icons/di";
import { FaBitcoin, FaThList } from "react-icons/fa";
import { FaHardDrive } from "react-icons/fa6";
import { LuTerminal } from "react-icons/lu";
import { MdMail } from "react-icons/md";
import { SiWireguard } from "react-icons/si";

export default function Home() {
  return (
    <div className="h-screen flex flex-col">
      <Navbar />
      <div className="flex flex-col gap-y-20 py-10 items-center justify-center flex-grow md:flex-row md:gap-x-20 ">

        <main className="flex items-center justify-start flex-col gap-y-10 px-10 md:px-0 md:gap-y-20">
          <div className="flex flex-col items-start justify-center gap-y-8">

            <div className="flex items-center justify-center gap-x-2 md:gap-x-5">
              <Image
                src="/assets/icon.png"
                alt="logo"
                width={70}
                height={70}
                className="inline-block"
              />
              <h1 className="text-4xl font-bold text-white md:text-7xl md:mb-2.5">
                config-lsp
              </h1>
            </div>

            <p className="text-2xl text-gray-400">
              The Language Server for your configuration files
            </p>
          </div>
          <div className="flex flex-col items-center justify-center gap-y-8">
            <div className="flex items-center justify-center gap-x-2">
              <Link
                href="https://marketplace.visualstudio.com/items?itemName=myzel394.config-lsp"
                className="px-6 py-3 font-bold text-white rounded-lg bg-primary hover:bg-gray-900 transition-colors cursor-pointer"
                rel="noopener noreferrer"
              >
                Download for VS Code
              </Link>
              <Link
                href="https://github.com/Myzel394/config-lsp"
                className="px-6 py-3 font-bold text-white rounded-lg cursor-pointer border-gray-800 border"
                rel="noopener noreferrer"
              >
                Check on GitHub
              </Link>
            </div>
            <div className="flex flex-col items-start justify-center gap-y-4">
              <p className="text-gray-500 text-lg">
                Supported languages:
              </p>
              <div className="grid grid-cols-3 gap-4">
                <SupportedConfig Icon={SiWireguard} name="WireGuard" />
                <SupportedConfig Icon={FaBitcoin} name="Bitcoin" />
                <SupportedConfig Icon={FaHardDrive} name="fstab" />
                <SupportedConfig Icon={DiTerminal} name="SSH" />
                <SupportedConfig Icon={LuTerminal} name="SSH Daemon" />
                <SupportedConfig Icon={MdMail} name="aliases" />
                <SupportedConfig Icon={FaThList} name="hosts" />
              </div>
            </div>
          </div>
        </main>

        <div className="w-full px-10 aspect-[1.1] rounded-[45px] overflow-hidden relative bg-[#1C1C1C] md:w-[40em]" id="preview-video">
          <Image
            src="/assets/preview-video_first-frame.jpg"
            alt=""
            sizes="100%"
            objectFit="cover"
            fill
            className="z-1"
          />
          <video width="100%" height="100%" autoPlay loop muted preload="auto" className="z-2 absolute top-0 left-0 object-cover">
            <source
              src="/assets/preview-video.mp4"
              type="video/mp4"
            />
            Your browser does not support the video tag.
          </video>
        </div>

      </div>
    </div>
  );
}

