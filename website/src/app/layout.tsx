import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "config-lsp",
  description: "The Language Server for your configuration files",
  keywords: [
    "config-lsp",
    "language server",
    "configuration files",
    "wireguard",
    "bitcoin",
    "fstab",
    "ssh",
    "ssh daemon",
    "aliases",
    "hosts",
  ],

  authors: [
    {
      name: "Myzel394",
      url: "https://github.com/Myzel394",
    },
  ],

  creator: "Myzel394",
  themeColor: "rgb(14,17,23)",
  robots: {
    index: true,
    follow: true,
  },

  twitter: {
    creator: "@Myzel394",
    title: "config-lsp",
    description: "The Language Server for your configuration files",
    card: "summary_large_image",
  },

  openGraph: {
    title: "config-lsp",
    description: "The Language Server for your configuration files",
    url: "https://myzel394.app",
    siteName: "config-lsp",
    type: "website",
    locale: "en_US",
  },

  icons: {
    icon: "/assets/icon.png",
    shortcut: "/assets/icon.png",
    apple: "/assets/icon.png",
  },


};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <html lang="en">
        <body
          className={`${geistSans.variable} ${geistMono.variable} antialiased`}
        >
          <div id="bg" />
          {children}
        </body>
      </html>
    </>
  );
}
