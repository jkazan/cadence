import { useState } from "react";
import { Button } from "@/components/ui/button";
import { SpotifyApi } from "@spotify/web-api-ts-sdk";
import { PlaylistBrowser } from "../src/PlaylistBrowser";
import { TempoFilter } from "../src/TempoFilter";
import { PlaylistGenerator } from "../src/PlaylistGenerator";
import { Music, Activity, Zap } from "lucide-react";

const Index = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [selectedPlaylists, setSelectedPlaylists] = useState([]);
  const [tempoRange, setTempoRange] = useState({ min: 120, max: 180 });

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      {/* Header */}
      <header className="border-b border-white/10 backdrop-blur-lg bg-black/20">
        <div className="container mx-auto px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-3">
              <div className="p-2 bg-gradient-to-r from-green-400 to-blue-500 rounded-lg">
                <Activity className="h-6 w-6 text-white" />
              </div>
              <h1 className="text-2xl font-bold text-white">RunBeat</h1>
            </div>
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-2 text-sm text-gray-300">
                <Zap className="h-4 w-4" />
                <span>Perfect Running Playlists</span>
              </div>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="container mx-auto px-6 py-8">
        {!isAuthenticated ? (
          <div className="max-w-2xl mx-auto text-center space-y-8">
            <div className="space-y-4">
              <h2 className="text-5xl font-bold text-white leading-tight">
                Create Perfect
                <span className="bg-gradient-to-r from-green-400 to-blue-500 bg-clip-text text-transparent">
                  {" "}Running{" "}
                </span>
                Playlists
              </h2>
              <p className="text-xl text-gray-300 leading-relaxed">
                Generate playlists with your favorite music, perfectly matched to your running cadence using BPM analysis.
              </p>
            </div>
            
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6 my-12">
              <div className="p-6 bg-white/5 backdrop-blur-lg rounded-xl border border-white/10">
                <Music className="h-8 w-8 text-green-400 mb-4 mx-auto" />
                <h3 className="text-lg font-semibold text-white mb-2">Your Music</h3>
                <p className="text-gray-400 text-sm">Browse your Spotify playlists and public collections</p>
              </div>
              <div className="p-6 bg-white/5 backdrop-blur-lg rounded-xl border border-white/10">
                <Activity className="h-8 w-8 text-blue-400 mb-4 mx-auto" />
                <h3 className="text-lg font-semibold text-white mb-2">BPM Analysis</h3>
                <p className="text-gray-400 text-sm">Analyze tempo of each track for perfect running rhythm</p>
              </div>
              <div className="p-6 bg-white/5 backdrop-blur-lg rounded-xl border border-white/10">
                <Zap className="h-8 w-8 text-purple-400 mb-4 mx-auto" />
                <h3 className="text-lg font-semibold text-white mb-2">Perfect Match</h3>
                <p className="text-gray-400 text-sm">Generate playlists within your ideal cadence range</p>
              </div>
            </div>

            <SpotifyAuth onAuthenticated={() => setIsAuthenticated(true)} />
          </div>
        ) : (
          <div className="space-y-8">
            <div className="text-center space-y-2">
              <h2 className="text-3xl font-bold text-white">Build Your Perfect Running Playlist</h2>
              <p className="text-gray-300">Select playlists and set your tempo preferences</p>
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
              <div className="lg:col-span-2">
                <PlaylistBrowser 
                  selectedPlaylists={selectedPlaylists}
                  onSelectionChange={setSelectedPlaylists}
                />
              </div>
              
              <div className="space-y-6">
                <TempoFilter 
                  tempoRange={tempoRange}
                  onRangeChange={setTempoRange}
                />
                
                <PlaylistGenerator 
                  selectedPlaylists={selectedPlaylists}
                  tempoRange={tempoRange}
                />
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
};

export default Index;
