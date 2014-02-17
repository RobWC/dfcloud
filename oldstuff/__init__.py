#!/usr/bin/python
import hashlib
import os
import gevent

class FileSet(object):
  def __init__(self,path):
    self.path = path
    self.files = []
  def fileScan(self):
    #find all the files in the specified path
    #add
    totalItems = 0
    totalSize = 0
    itemsInPath = os.listdir(self.path)
    for x in itemsInPath:
      pathToItem = self.path + '/' + x
      if os.path.isdir(pathToItem):
        #recurse
        for (path, dirs, files) in os.walk(pathToItem):
          for filename in files:
            pathToFile = path + '/' + filename
            testfile = File(pathToFile)
            self.files.append(testfile)
            testfile.fileChecksum()
            totalSize = totalSize + testfile.checksum['size']
            totalItems = totalItems + 1;

      elif os.path.isfile(pathToItem):
        testfile = File(pathToItem)
        self.files.append(testfile)
        testfile.fileChecksum()
        totalSize = totalSize + testfile.checksum['size']
        totalItems = totalItems + 1;


    print 'Total File Count: {0}\nTotal Size of Files: {1}\n'.format(totalItems,totalSize)


class File(object):
  def __init__(self,path):
    self.path = os.path.abspath(path)
    self.name = 'foo'
    self.fileSize = os.path.getsize(self.path)
    self.checksum = {"value":None,"type":"sha256","size":self.fileSize}
  def readFile(self):
    return open(self.path).read()
  def fileChecksum(self):
    hasher = hashlib.sha256()
    fileHandle = open(self.path,'rb')
    data = fileHandle.read(1024000)
    dataLen = 0
    while len(data) > 0:
      hasher.update(data)
      dataLen = dataLen + len(data)
      data = fileHandle.read(102400)
    digext = hasher.hexdigest()
    print self.path
    print digext
    print dataLen
    self.checksum['value'] = digext

class DFHerod(object):
  def __init__(self):
    pass
  def fileChecksum(self,file):
    hasher = hashlib.sha256()
    fileHandle = file.readFile()
    print hasher.hexdigest()

fileTestSet = FileSet('/home/rcameron/code/dfcloud/save')
fileTestSet.fileScan()
