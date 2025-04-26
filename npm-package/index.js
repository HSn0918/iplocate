#!/usr/bin/env node

/**
 * iplocate-mcp - IP地址和经纬度位置查询MCP服务
 *
 * 此脚本会使用当前目录下的可执行文件启动iplocate的MCP服务
 */

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');
const os = require('os');

// 检测操作系统和架构
const platform = os.platform();
const arch = os.arch();

console.log(`操作系统: ${platform}, 架构: ${arch}`);
console.log('正在启动 IP 位置查询 MCP 服务...');

// 确定可执行文件路径
let executableName;
if (platform === 'win32') {
  executableName = 'iplocate.exe';
} else if (platform === 'darwin') {
  executableName = 'iplocate-darwin';
} else if (platform === 'linux') {
  if (arch === 'x64') {
    executableName = 'iplocate-linux';
  } else if (arch === 'arm64') {
    executableName = 'iplocate-linux-arm64';
  } else {
    console.error(`不支持的Linux架构: ${arch}`);
    process.exit(1);
  }
} else {
  console.error(`不支持的操作系统: ${platform}`);
  process.exit(1);
}

// 可执行文件相对路径
const relativePath = path.join('bin', executableName);
// 可执行文件的绝对路径
const executablePath = path.resolve(__dirname, relativePath);

console.log(`使用可执行文件: ${executablePath}`);

// 确保可执行文件存在
try {
  fs.accessSync(executablePath, fs.constants.F_OK);
} catch (error) {
  console.error(`错误: 找不到可执行文件 ${executablePath}`);
  console.error('请确保二进制文件已放置在正确的位置');
  process.exit(1);
}

// 确保文件有可执行权限
try {
  fs.accessSync(executablePath, fs.constants.X_OK);
} catch (error) {
  console.log('设置可执行权限...');
  try {
    // 在非Windows平台上设置可执行权限
    if (platform !== 'win32') {
      fs.chmodSync(executablePath, 0o755);
    }
  } catch (chmodErr) {
    console.error(`无法设置可执行权限: ${chmodErr.message}`);
    process.exit(1);
  }
}

// 启动 MCP 服务
console.log('启动 iplocate MCP 服务...');

// 设置进程环境变量
const env = { ...process.env };

// 启动 iplocate MCP 服务
const iplocateProcess = spawn(executablePath, ['mcp'], {
  stdio: 'inherit',
  env: env
});

// 处理进程退出
iplocateProcess.on('close', (code) => {
  console.log(`iplocate MCP 服务已退出，退出码 ${code}`);
});

// 处理进程错误
iplocateProcess.on('error', (err) => {
  console.error(`启动 iplocate MCP 服务出错: ${err.message}`);
  process.exit(1);
});

// 处理进程终止信号
process.on('SIGINT', () => {
  console.log('接收到 SIGINT 信号，正在关闭服务...');
  iplocateProcess.kill('SIGINT');
});

process.on('SIGTERM', () => {
  console.log('接收到 SIGTERM 信号，正在关闭服务...');
  iplocateProcess.kill('SIGTERM');
});

console.log('服务已启动，等待请求...');