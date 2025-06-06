import { useState } from 'react';

export default function BlackoutInterface() {
  const [activeTab, setActiveTab] = useState('login');
  const [loginForm, setLoginForm] = useState({ username: '', password: '' });
  const [registerForm, setRegisterForm] = useState({ username: '', email: '', password: '', confirmPassword: '' });
  const [gameForm, setGameForm] = useState({ gameName: '' });

  const handleLogin = () => {
    console.log('Logging in:', loginForm);
  };

  const handleRegister = () => {
    console.log('Registering:', registerForm);
  };

  const handleCreateGame = () => {
    console.log('Creating game:', gameForm);
  };

  return (
    <div className="min-h-screen bg-stone-700 relative">
      {/* Dingy concrete wall texture */}
      <div className="absolute inset-0 bg-gradient-to-br from-stone-600 via-stone-700 to-stone-800"></div>
      
      {/* Water stains and discoloration */}
      <div className="absolute inset-0" style={{
        backgroundImage: `
          radial-gradient(ellipse at 20% 30%, rgba(120, 113, 108, 0.4) 0%, transparent 60%),
          radial-gradient(ellipse at 80% 10%, rgba(87, 83, 81, 0.3) 0%, transparent 50%),
          radial-gradient(ellipse at 60% 80%, rgba(168, 162, 158, 0.2) 0%, transparent 40%)
        `,
        backgroundSize: '400px 300px, 300px 200px, 500px 400px'
      }}></div>
      
      {/* Flickering fluorescent light effect */}
      <div className="absolute top-0 left-1/4 right-1/4 h-2 bg-gradient-to-r from-transparent via-stone-300 to-transparent opacity-30"></div>
      
      {/* Cigarette smoke trails */}
      <div className="absolute top-1/3 left-1/4 w-px h-16 bg-gradient-to-t from-stone-400 to-transparent opacity-40 blur-sm"></div>
      <div className="absolute top-2/3 right-1/3 w-px h-12 bg-gradient-to-t from-stone-500 to-transparent opacity-30 blur-sm"></div>
      <div className="absolute bottom-1/4 left-2/3 w-px h-8 bg-gradient-to-t from-stone-400 to-transparent opacity-35 blur-sm"></div>
      
      <div className="relative z-10 container mx-auto px-6 py-12">
        {/* Weathered signage */}
        <div className="text-center mb-12">
          <div className="inline-block bg-stone-800 px-8 py-4 border-2 border-stone-600 shadow-lg">
            <h1 className="text-4xl font-black text-stone-200 tracking-wider" 
                style={{ fontFamily: 'monospace' }}>
              BLACKOUT
            </h1>
          </div>
        </div>

        {/* Weathered tab bar */}
        <div className="flex justify-center mb-8">
          <div className="bg-stone-800 border-2 border-stone-600 shadow-inner">
            {['login', 'register', 'create'].map((tab) => (
              <button
                key={tab}
                onClick={() => setActiveTab(tab)}
                className={`px-6 py-3 text-xs uppercase tracking-wider font-mono border-r-2 border-stone-600 last:border-r-0 transition-colors ${
                  activeTab === tab
                    ? 'bg-stone-600 text-stone-100'
                    : 'text-stone-400 hover:text-stone-200 hover:bg-stone-700'
                }`}
              >
                {tab === 'create' ? 'NEW GAME' : tab}
              </button>
            ))}
          </div>
        </div>

        {/* Main form - like a cutout in the wall */}
        <div className="max-w-sm mx-auto">
          <div className="bg-stone-800 border-2 border-stone-600 shadow-2xl relative">
            {/* Worn edges and scratches */}
            <div className="absolute top-0 left-4 w-8 h-px bg-stone-500 opacity-60"></div>
            <div className="absolute bottom-2 right-6 w-4 h-px bg-stone-500 opacity-40"></div>
            <div className="absolute left-0 top-8 w-px h-6 bg-stone-500 opacity-50"></div>
            
            <div className="p-6">
              {/* Login */}
              {activeTab === 'login' && (
                <div>
                  <h2 className="text-lg font-black text-stone-200 mb-6 uppercase tracking-wide font-mono">
                    ENTER
                  </h2>
                  <div className="space-y-4">
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Name
                      </label>
                      <input
                        type="text"
                        value={loginForm.username}
                        onChange={(e) => setLoginForm({...loginForm, username: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Password
                      </label>
                      <input
                        type="password"
                        value={loginForm.password}
                        onChange={(e) => setLoginForm({...loginForm, password: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <button
                      onClick={handleLogin}
                      className="w-full bg-stone-600 hover:bg-stone-500 text-stone-100 font-black py-3 uppercase tracking-wider text-sm transition-colors font-mono border border-stone-500"
                    >
                      IN
                    </button>
                  </div>
                </div>
              )}

              {/* Register */}
              {activeTab === 'register' && (
                <div>
                  <h2 className="text-lg font-black text-stone-200 mb-6 uppercase tracking-wide font-mono">
                    SIGN UP
                  </h2>
                  <div className="space-y-3">
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Name
                      </label>
                      <input
                        type="text"
                        value={registerForm.username}
                        onChange={(e) => setRegisterForm({...registerForm, username: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Email
                      </label>
                      <input
                        type="email"
                        value={registerForm.email}
                        onChange={(e) => setRegisterForm({...registerForm, email: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Password
                      </label>
                      <input
                        type="password"
                        value={registerForm.password}
                        onChange={(e) => setRegisterForm({...registerForm, password: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Again
                      </label>
                      <input
                        type="password"
                        value={registerForm.confirmPassword}
                        onChange={(e) => setRegisterForm({...registerForm, confirmPassword: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <button
                      onClick={handleRegister}
                      className="w-full bg-stone-600 hover:bg-stone-500 text-stone-100 font-black py-3 uppercase tracking-wider text-sm transition-colors font-mono border border-stone-500"
                    >
                      JOIN
                    </button>
                  </div>
                </div>
              )}

              {/* Create Game */}
              {activeTab === 'create' && (
                <div>
                  <h2 className="text-lg font-black text-stone-200 mb-6 uppercase tracking-wide font-mono">
                    NEW GAME
                  </h2>
                  <div className="space-y-4">
                    <div>
                      <label className="block text-stone-400 text-xs uppercase tracking-wider mb-1 font-mono">
                        Game Name
                      </label>
                      <input
                        type="text"
                        value={gameForm.gameName}
                        onChange={(e) => setGameForm({...gameForm, gameName: e.target.value})}
                        className="w-full px-3 py-2 bg-stone-700 text-stone-200 border border-stone-600 focus:border-stone-500 focus:outline-none font-mono text-sm"
                      />
                    </div>
                    <button
                      onClick={handleCreateGame}
                      className="w-full bg-stone-600 hover:bg-stone-500 text-stone-100 font-black py-3 uppercase tracking-wider text-sm transition-colors font-mono border border-stone-500"
                    >
                      START
                    </button>
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}