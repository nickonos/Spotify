import { ReactNode } from 'react'
import { ResizableHandle, ResizablePanel, ResizablePanelGroup } from './components/ui/resizable'
import { HomeIcon, SearchIcon } from './components/icons';

function App() {

  return (
    <body className="dark h-screen">

      <Layout>hello</Layout>
    </body>
  )
}

export default App


const Layout = ({children}: {children: ReactNode}) => (
  <ResizablePanelGroup
    direction={"horizontal"}
    className="flex h-screen w-screen gap-1 p-2"
  >
    <ResizablePanel
      defaultSize={20}
      maxSize={40}
      className="h-full bg-black p-1"
    >
      <ResizablePanelGroup direction="vertical" className="flex gap-1">
        <ResizablePanel defaultSize={10}>
          <div className="flex">
            <HomeIcon />
            <div>Home</div>
          </div>
          <div className="flex">
            <SearchIcon />
            <div>Search</div>
          </div>
        </ResizablePanel>
        <ResizableHandle hidden />
        <ResizablePanel></ResizablePanel>
      </ResizablePanelGroup>
    </ResizablePanel>
    <ResizableHandle hidden />
    <ResizablePanel>
      {children}
    </ResizablePanel>
  </ResizablePanelGroup>
);